#!/bin/sh

go run main.go https://golang.org target

cd find
go test -v
