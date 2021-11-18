#!/bin/sh

echo "tag:"
curl -s https://www.w3.org/TR/2006/REC-xml11-20060816/ | go run main.go div div h2

echo "\nid:"
curl -s https://www.w3.org/TR/2006/REC-xml11-20060816/ | go run main.go '#dt-app'

echo "\nclass:"
curl -s https://www.w3.org/TR/2006/REC-xml11-20060816/ | go run main.go .back .div1 h2
