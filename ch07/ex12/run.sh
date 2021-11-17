#!/bin/sh

set -mx # enable job control

go run main.go &
sleep 1

curl 'http://localhost:8000/list'
curl 'http://localhost:8000/price?item=shoes'

curl 'http://localhost:8000/create?item=book&price=3'
curl 'http://localhost:8000/create?item=shoes&price=7'
curl 'http://localhost:8000/create?item=cloth&price=x'

curl 'http://localhost:8000/update?item=socks&price=6'
curl 'http://localhost:8000/update?item=desk&price=20'
curl 'http://localhost:8000/update?item=shoes&price=x'

curl 'http://localhost:8000/delete?item=shoes'
curl 'http://localhost:8000/delete?item=chair'

curl 'http://localhost:8000/list'

kill %1
