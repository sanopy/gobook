#!/bin/sh

set -m # enable job control

go run mandelbrot.go &
sleep 1

curl 'http://localhost:8000' > default.png
curl 'http://localhost:8000?scale=2' > 2x.png
curl 'http://localhost:8000/?scale=4&y=-1' > vertical_shift.png
curl 'http://localhost:8000/?scale=4&x=-1' > horizontal_shift.png
curl 'http://localhost:8000/?scale=4&y=.5&x=-1' > shift.png

kill %1
