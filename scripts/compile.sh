#!/bin/bash
# Compiles the main app for all platforms into the build directory.
echo 'Compiling into ./build'
mkdir -p build
for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        docker-compose exec cli env GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o ./build/main-${GOOS}-${GOARCH} ./cmd/beetroot/main.go
    done
done
chmod +x ./build/*
echo 'Done.'
