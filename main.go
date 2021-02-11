package main

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// https://stackoverflow.com/a/46735876
func getCertificatesPEM(address string) ([]byte, error) {
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
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

func updateFedora(certs []byte) {
	err := ioutil.WriteFile("/etc/pki/ca-trust/source/anchors/insecure.pem", certs, 0644)
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("update-ca-trust")

	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	uriPtr := flag.String("uri", "jmainguy.com:443", "A hostname and port, jmainguy.com:443 for example")
	pemPtr := flag.String("pem", "insecure.pem", "pem file to write to, insecure.pem by default")
	updateFedoraPtr := flag.Bool("updateFedora", false, "write pem to /etc/pki/ca-trust/source/anchors and run update-ca-trust")

	flag.Parse()

	certs, err := getCertificatesPEM(*uriPtr)
	if err != nil {
		fmt.Println(err)
	}
	if *updateFedoraPtr {
		updateFedora(certs)
	} else {
		err = ioutil.WriteFile(*pemPtr, certs, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}
}
