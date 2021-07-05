# certificatedownloader
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/certificatedownloader)](https://goreportcard.com/badge/github.com/Jmainguy/certificatedownloader)
[![Release](https://img.shields.io/github/release/Jmainguy/certificatedownloader.svg?style=flat-square)](https://github.com/Jmainguy/certificatedownloader/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/certificatedownloader/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/certificatedownloader?branch=main)

This tool is designed to download a certificate from a remote source.

It also has support for updating your OS to trust the cert (currently only Fedora, however feel free to contribute back code for your OS of choice)

I use this tool to trust domains ending .local and other untrusted certificates.

## Usage
```/bin/bash
Usage of certificatedownloader:
  -pem string
    	pem file to write to, same name as host.port.pem by default, as interpretted from uri
  -timeout int
    	Timeout in seconds (default 10)
  -updateFedora
    	write pem to /etc/pki/ca-trust/source/anchors and run update-ca-trust
  -uri string
    	A hostname and port, jmainguy.com:443 for example

  To save jmainguy.com:443 cert, run certificatedownloader --uri jmainguy.com:443 for example
```

## Examples

```/bin/bash
# Download cert from jmainguy.com on port 443, and tell OS to trust it
certificatedownloader --uri jmainguy.com:443 --updateFedora

# Attempt to download cert from google.com on port 4444, but timeout after 5 seconds
certificatedownloader --uri google.com:4444

# Download certificate to /tmp/example.pem
certificatedownloader --uri jmainguy.com:443 --pem /tmp/example.pem
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/certificatedownloader/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```
