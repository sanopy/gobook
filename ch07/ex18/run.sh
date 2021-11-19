#!/bin/sh

curl -s https://www.w3.org/TR/2006/REC-xml11-20060816/ | go run main.go

cd xmlnode
go test -v
