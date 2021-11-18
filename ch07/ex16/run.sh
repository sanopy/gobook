#!/bin/sh

set -mx # enable job control

go run main.go &
sleep 1

# 1 + 2
curl 'http://localhost:8000/?expr=1%2B2'

# 1 + 2 * 3
curl 'http://localhost:8000/?expr=1%2B2%2A3'

# x + y, x = 5, y = 3
curl 'http://localhost:8000/?expr=x%2By&x=5&y=3'

# x < y ? y - x : x - y, x = 5, y = 10
curl 'http://localhost:8000/?expr=x+<+y+%3F+y+-+x+%3A+x+-+y&x=5&y=10'

kill %1
