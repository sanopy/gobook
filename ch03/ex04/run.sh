#!/bin/sh

set -m # enable job control

go run surface.go &
sleep 1

curl http://localhost:8000 > default.svg
curl http://localhost:8000?f=egg > egg.svg
curl http://localhost:8000?f=mogul > mogul.svg
curl http://localhost:8000?f=saddle > saddle.svg

kill %1