#!/bin/sh

if [ ! -f .env ]; then
  echo '.env is not found. Run following commands.'
  echo 'echo "APIKEY=${Your API Key}" > .env'
  exit
fi

export $(cat .env | xargs)

go run main.go titanic
go run main.go imitation game
