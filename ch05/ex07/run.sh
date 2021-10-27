#!/bin/sh

go run main.go https://golang.org

cd pretty
go test -v
