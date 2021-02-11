# certificateDownloader
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/certificateDownloader)](https://goreportcard.com/badge/github.com/Jmainguy/certificateDownloader)
[![Release](https://img.shields.io/github/release/Jmainguy/certificateDownloader.svg?style=flat-square)](https://github.com/Jmainguy/certificateDownloader/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/certificateDownloader/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/certificateDownloader?branch=main)

This command is designed to to bump the version in CHANGELOG.md and RELEASE

## Usage
```/bin/bash
Usage of ./certificateDownloader:
  -pem string
    	pem file to write to, insecure.pem by default (default "insecure.pem")
  -updateFedora
    	write pem to /etc/pki/ca-trust/source/anchors and run update-ca-trust
  -uri string
    	A hostname and port, jmainguy.com:443 for example (default "jmainguy.com:443")
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
