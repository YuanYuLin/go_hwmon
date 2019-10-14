#!/bin/bash
echo "PATH="$PATH
echo "GOARCH="$GOARCH
echo "GOROOT="$GOROOT
echo "GOOS="$GOOS
echo "GOPATH="$GOPATH
$GOROOT/bin/go build
