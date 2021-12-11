#!/bin/bash

set -m # enable job control

function final {
  kill -TERM %1
}

trap "final; exit 1" SIGINT # trap ctrl+c

go run reverb1/reverb.go &

sleep 1

go run netcat.go

final
