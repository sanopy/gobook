#!/bin/bash

set -x

go test -bench=.

time go run mandelbrot.go > out.png
time go run mandelbrot.go -parallel > out.png
