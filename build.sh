#!/bin/bash

pushd ui > /dev/null
npm ci
npm run build
popd > /dev/null

rm -rf src/ui
mv ui/dist src/ui

pushd src > /dev/null
GOOS=darwin GOARCH=amd64 go build -trimpath -a -ldflags="-w -s" -o ../golang-embed-demo-mac
GOOS=windows GOARCH=amd64 go build -trimpath -a -ldflags="-w -s" -o ../golang-embed-demo-win.exe
popd > /dev/null
