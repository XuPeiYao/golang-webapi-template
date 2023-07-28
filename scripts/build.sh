#!/bin/bash

rm -f -r ./dist
mkdir -p ./dist

go build -o ./dist ./cmd/api-server 

cp -r ./assets ./dist
cp -r ./configs ./dist
