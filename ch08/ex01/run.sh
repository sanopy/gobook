#!/bin/bash

set -m # enable job control

function final {
  kill -TERM %3
  kill -TERM %2
  kill -TERM %1
}

trap "final; exit 1" SIGINT # trap ctrl+c

TZ=US/Eastern go run clock/clock.go -port 8010 &
TZ=Asia/Tokyo go run clock/clock.go -port 8020 &
TZ=Europe/London go run clock/clock.go -port 8030 &

sleep 1

go build -o clockwall -race clockwall.go
./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030 # go run では trap に失敗
