#!/bin/bash

if [ -f ~/go/extra/golang-crosscompile/crosscompile.bash ]; then
    . ~/go/extra/golang-crosscompile/crosscompile.bash
fi

if type -t go-build-all | grep -i "function"; then
    go-build-all
    ls -1 -- -* | xargs -n1 -i{} mv -- "{}" "exodin{}"
    ls -1 -- exodin-* | xargs -n1 -i{} upx "{}"
    goupx exodin-linux-amd64
else
	exit
    GOOS=windows GOARCH=386 go build
    mv exodin.exe exodin-x86.exe
    upx exodin-x86.exe

    GOOS=windows GOARCH=amd64 go build
    mv exodin.exe exodin-amd64.exe
    upx exodin-amd64.exe

    GOARCH=386 go build
    mv exodin exodin-x86
    upx exodin-x86

    GOARCH=amd64 go build
    mv exodin exodin-amd64
    goupx exodin-amd64
fi
