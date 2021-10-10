#!/bin/sh

echo 'x' | go run hash.go
echo 'x' | go run hash.go -algo sha256
echo 'x' | go run hash.go -algo sha384
echo 'x' | go run hash.go -algo sha512
