#!/bin/sh

set -m # enable job control

go run main.go &
sleep 1

curl 'http://localhost:8000/?repo=golang/go'

kill %1
