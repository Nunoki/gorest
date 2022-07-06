#!/bin/bash
OUTPUT_DIR=./build

HELPTEXT="Compile the executables for all platforms into the $OUTPUT_DIR directory."
source $(dirname "$0")/_help_text.sh $@

echo "Compiling into $OUTPUT_DIR"
mkdir -p $OUTPUT_DIR
for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
        GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o $OUTPUT_DIR/beetroot-${GOOS}-${GOARCH} ./cmd/beetroot/main.go
        GOOS=${GOOS} GOARCH=${GOARCH} go build -v -o $OUTPUT_DIR/migrate-${GOOS}-${GOARCH} ./cmd/migrate/main.go
    done
done
chmod +x $OUTPUT_DIR/*
echo 'Done.'
