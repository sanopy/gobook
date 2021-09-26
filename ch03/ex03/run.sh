#!/bin/sh

go run surface.go > default.svg
go run surface.go -f egg > egg.svg
go run surface.go -f mogul > mogul.svg
go run surface.go -f saddle > saddle.svg
