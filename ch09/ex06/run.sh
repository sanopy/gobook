#!/bin/bash

set -x

if [ $(nproc) -ge 1 ]; then
  GOMAXPROCS=1 go test -bench=.
fi

if [ $(nproc) -ge 2 ]; then
  GOMAXPROCS=2 go test -bench=.
fi

if [ $(nproc) -ge 4 ]; then
  GOMAXPROCS=4 go test -bench=.
fi

if [ $(nproc) -ge 6 ]; then
  GOMAXPROCS=6 go test -bench=.
fi

if [ $(nproc) -ge 8 ]; then
  GOMAXPROCS=8 go test -bench=.
fi

if [ $(nproc) -ge 12 ]; then
  GOMAXPROCS=12 go test -bench=.
fi

if [ $(nproc) -ge 16 ]; then
  GOMAXPROCS=16 go test -bench=.
fi
