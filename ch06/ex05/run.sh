#!/bin/sh

go test -v

# for linux; `GOARCH=386` is not supported for darwin
if [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
  # 64bit
  GOARCH=amd64 go test -v

  # 32bit
  GOARCH=386 go test -v
fi
