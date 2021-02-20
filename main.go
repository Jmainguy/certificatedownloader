package main

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

// https://stackoverflow.com/a/46735876
func getCertificatesPEM(dialTimeout int, address string) ([]byte, error) {
	dialer := &net.Dialer{
		Timeout: time.Duration(dialTimeout) * time.Second,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		if strings.Contains(err.Error(), "i/o timeout") {
			fmt.Printf("Timed out connecting to %s, waited %d seconds\n", address, dialTimeout)
			fmt.Println("if you are confident that you can connect, consider increasing the timeout with --timeout 30, for 30 seconds")
			return []byte(""), err
		}
		return []byte(""), err
	}
	defer conn.Close()
	var b bytes.Buffer
	for _, cert := range conn.ConnectionState().PeerCertificates {
		err := pem.Encode(&b, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		})
		if err != nil {
			return []byte(""), err
		}
	}
	certs := b.Bytes()
	return certs, nil
}

func updateFedora(certs []byte, pemName string) {
	path := "/etc/pki/ca-trust/source/anchors/" + pemName
	err := ioutil.WriteFile(path, certs, 0644)
	if err != nil {
		if strings.Contains(err.Error(), "permission denied") {
			fmt.Println(err)
			fmt.Println("  Try running with sudo, sudo !!")
			os.Exit(1)
		}
		fmt.Println(err)
		os.Exit(1)
	}
	cmd := exec.Command("update-ca-trust")

	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	uriPtr := flag.String("uri", "", "A hostname and port, jmainguy.com:443 for example")
	pemPtr := flag.String("pem", "", "pem file to write to, insecure.pem by default")
	updateFedoraPtr := flag.Bool("updateFedora", false, "write pem to /etc/pki/ca-trust/source/anchors and run update-ca-trust")
	timeoutPtr := flag.Int("timeout", 10, "Timeout in seconds")

	flag.Parse()

	if *uriPtr == "" {
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("  To save jmainguy.com:443 cert, run certificateDownloader --uri jmainguy.com:443 for example")
		fmt.Println("")
		os.Exit(1)
	}

	// Replace https:// and / with "", this will allow a host like https://jmainguy.com/ to be passed
	// Even though, you really shouldnt be passing it in this format
	uri := strings.ReplaceAll(*uriPtr, "https://", "")
	uri = strings.ReplaceAll(uri, "/", "")

	uriArray := strings.Split(uri, ":")
	if len(uriArray) == 1 {
		uri = uri + ":443"
	} else if len(uriArray) > 2 {
		fmt.Printf("Please format uri as host:port, you provided %s\n", uri)
	}
	certs, err := getCertificatesPEM(*timeoutPtr, uri)
	var pemName string
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *pemPtr == "" {
		pemName = strings.ReplaceAll(uri, ":", ".")
	} else {
		pemName = *pemPtr
	}
	pemName = pemName + ".pem"
	if *updateFedoraPtr {
		updateFedora(certs, pemName)
	} else {
		err = ioutil.WriteFile(pemName, certs, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
