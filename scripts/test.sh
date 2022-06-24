#!/bin/bash
# Shorthand for a simple go test call, but with colorize output
OUTPUT=$(go test ./...)
source "$(dirname "$0")/_colorized_test_results.sh"
