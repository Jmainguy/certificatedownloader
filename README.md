# certificateDownloader
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/certificateDownloader)](https://goreportcard.com/badge/github.com/Jmainguy/certificateDownloader)
[![Release](https://img.shields.io/github/release/Jmainguy/certificateDownloader.svg?style=flat-square)](https://github.com/Jmainguy/certificateDownloader/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/certificateDownloader/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/certificateDownloader?branch=main)

This tool is designed to download a certificate from a remote source.

It also has support for updating your OS to trust the cert (currently only Fedora, however feel free to contribute back code for your OS of choice)

I use this tool to trust domains ending .local and other untrusted certificates.

## Usage
```/bin/bash
Usage of certificateDownloader:
  -pem string
    	pem file to write to, same name as host.port.pem by default, as interpretted from uri
  -timeout int
    	Timeout in seconds (default 10)
  -updateFedora
    	write pem to /etc/pki/ca-trust/source/anchors and run update-ca-trust
  -uri string
    	A hostname and port, jmainguy.com:443 for example

  To save jmainguy.com:443 cert, run certificateDownloader --uri jmainguy.com:443 for example
```
```/bin/bash
./certificateDownloader -updateFedora
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/certificateDownloader/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```
