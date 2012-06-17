#!/bin/bash

GOOS=windows GOARCH=386 go build
upx exodin.exe

GOARCH=386 go build
upx exodin
