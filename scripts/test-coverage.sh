#!/bin/bash
# Runs tests and generates test coverage results in the .test-coverage directory, then tells the
# system to try to open them directly.
OUTPUT_FILE=./.test-coverage/test-coverage.html
mkdir -p .test-coverage

echo 'Running tests'

OUTPUT=$(go test -coverprofile=.test-coverage/c.out -coverpkg=./...  ./... \
&& go tool cover -html=.test-coverage/c.out -o $OUTPUT_FILE)

STATUS=$?

source "$(dirname "$0")/_colorized_test_results.sh"

if [[ STATUS -ne 0 ]]; then 
    # if an error was returned, we exit early
    exit 1
fi

echo ""
echo "Test coverage report in HTML format generated in $OUTPUT_FILE"

if command -v open 1> /dev/null
then
    open $OUTPUT_FILE
elif command -v xdg-open 1> /dev/null
then
    xdg-open $OUTPUT_FILE 2> /dev/null
fi
