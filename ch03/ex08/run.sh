#!/bin/sh

go run mandelbrot.go -t complex64 > ./assets/complex64.png
go run mandelbrot.go -t complex128 > ./assets/complex128.png
go run mandelbrot.go -t float > ./assets/float.png
timeout 30 go run mandelbrot.go -t rat > /dev/null
