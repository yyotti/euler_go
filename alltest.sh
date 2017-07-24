#!/usr/bin/env bash
cd "$(dirname "${BASH_SOURCE:-$0}")"
echo GOPATH "$GOPATH"
echo GOROOT "$GOROOT"
find ./src -mindepth 1 -maxdepth 1 -type d -print0 | xargs -0 go test
