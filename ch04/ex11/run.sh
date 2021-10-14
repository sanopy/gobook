#!/bin/sh

if [ ! -f .env ]; then
  echo '.env is not found. Run following commands.'
  echo 'echo "GITHUB_USERNAME=${Your GitHub username}" > .env'
  echo 'echo "GITHUB_TOKEN=${Your GitHub Token}" >> .env'
  exit
fi

export $(cat .env | xargs)

echo '-------------------- get --------------------'
go run main.go get -repo "sanopy/gobook_test" -num 3

echo '-------------------- create --------------------'
go run main.go create -repo "sanopy/gobook_test" -title "test" -body "This is test."

echo '-------------------- create (editor) --------------------'
go run main.go create -repo "sanopy/gobook_test"

echo '-------------------- update (editor) --------------------'
go run main.go update -repo "sanopy/gobook_test" -num 5

echo '-------------------- close --------------------'
go run main.go close -repo "sanopy/gobook_test" -num 3