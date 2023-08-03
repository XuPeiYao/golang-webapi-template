#!/bin/bash

rm -f -r ./dist/api-server
mkdir -p ./dist
mkdir -p ./dist/api-server

go build -ldflags "-s -w" -o ./dist/api-server ./cmd/api-server

cp -r ./assets ./dist/api-server
cp -r ./configs ./dist/api-server
