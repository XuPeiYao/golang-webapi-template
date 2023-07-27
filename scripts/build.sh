#!/bin/bash

rm -f -r ./dist
mkdir -p ./dist
wire
sh ./scripts/update-swagger.sh
go mod download

if [ "$(expr substr $(uname -s) 1 5)" == "MINGW" ]; then
    go build -o ./dist/$(basename "$PWD").exe
else
    go build -o ./dist/$(basename "$PWD")
fi

cp -r ./assets ./dist
cp -r ./configs ./dist
