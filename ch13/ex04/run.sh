#!/bin/sh

set -x

wc -c /usr/share/dict/words
sha256sum < /usr/share/dict/words
go run main.go < /usr/share/dict/words | wc -c
go run main.go < /usr/share/dict/words | bunzip2 | sha256sum

cd bzip
go test -v
