#!/bin/bash
set -e

pushd ui > /dev/null
npm ci
npm run build
popd > /dev/null

rm -rf src/ui
mv ui/build src/ui

pushd src > /dev/null
GOOS=darwin GOARCH=arm64 go build -trimpath -a -ldflags="-w -s" -o ../golang-embed-demo-darwin-arm64
GOOS=windows GOARCH=amd64 go build -trimpath -a -ldflags="-w -s" -o ../golang-embed-demo-windows-amd64.exe
popd > /dev/null
