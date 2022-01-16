#!/bin/sh

cd `dirname $0`

cd ../../

tar cvf test.tar .
zip -r test.zip .

cd -
mv ../..//test.tar .
mv ../../test.zip .

mkdir test
cd test

go run ../sample/main.go < ../test.tar
go run ../sample/main.go < ../test.zip
