#!/bin/bash
echo "building index.wasm..."
GOARCH=wasm GOOS=js go build -o index.wasm main.go
echo "end"
