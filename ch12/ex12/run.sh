#!/bin/sh

set -mx # enable job control

go run main.go &
sleep 1

# email: test
curl 'http://localhost:12345/search?e=test'

# email: test@example.com
curl 'http://localhost:12345/search?e=test@example.com'

# credit: 1234567
curl 'http://localhost:12345/search?c=1234567'

# credit: 4111111111111111
curl 'http://localhost:12345/search?c=4111111111111111'

# zipcode: 1234567
curl 'http://localhost:12345/search?z=1234567'

# zipcode: 60601
curl 'http://localhost:12345/search?z=60601'

kill %1
