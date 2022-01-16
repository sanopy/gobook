#!/bin/sh

go build ../../ch03/ex05/mandelbrot.go

./mandelbrot | go run main.go -gif > mandelbrot.gif
./mandelbrot | go run main.go -jpeg > mandelbrot.jpg
./mandelbrot | go run main.go -png > mandelbrot.png
