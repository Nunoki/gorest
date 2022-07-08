#!/bin/bash
OUTPUT_DIR=./build

HELPTEXT="Compile the executables for all target platforms into the $OUTPUT_DIR directory. To see supported target platforms, run \`go tool dist list.\`"
source $(dirname "$0")/_help_text.sh $@

# attempt to get the list of supported target platforms
DIST_LIST=$(go tool dist list)
if [[ ! "$?" == "0" ]]; then
    echo "Couldn't get list of target platforms. Is go installed?"
    exit 1
fi
DIST_COUNT=$(echo "$DIST_LIST" | wc -l)

# ask for confirmation before building because it might take a long time
echo "Will attempt to build for $DIST_COUNT architectures. Continue? (y/n)"
read REPLY
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Nothing built."
    exit 0
fi

echo "Compiling into $OUTPUT_DIR"
mkdir -p $OUTPUT_DIR
COUNT_OK=0
COUNT_FAIL=0
for DIST in $(echo "$DIST_LIST"); do
    IFS=/ read -ra RES <<< "$DIST"
    echo "Building ${RES[0]}/${RES[1]}"

    GOOS=${RES[0]} GOARCH=${RES[1]} go build -v -o $OUTPUT_DIR/gorest-${RES[0]}-${RES[1]} ./cmd/gorest/main.go
    GOOS=${RES[0]} GOARCH=${RES[1]} go build -v -o $OUTPUT_DIR/migrate-${RES[0]}-${RES[1]} ./cmd/migrate/main.go

    # count results
    if [[ "$?" == "0" ]]; then
        COUNT_OK=$((COUNT_OK + 1))
    else
        COUNT_FAIL=$((COUNT_FAIL + 1))
    fi
done

chmod +x $OUTPUT_DIR/*
echo "Done. ${COUNT_OK} succeeded, ${COUNT_FAIL} failed."
