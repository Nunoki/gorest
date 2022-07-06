#!/bin/bash
# Helper script that outputs the contents of the $OUTPUT variable, which is assumed to be test
# results, manipulating it in order to colorize the labels "ok" into green and "FAIL" into red.
OUTPUT=$(echo "$OUTPUT" | sed 's/^ok/\\033[00;32mok\\033[0m/')
OUTPUT=$(echo "$OUTPUT" | sed 's/^FAIL/\\033[00;31mFAIL\\033[0m/')
echo -e "$OUTPUT"
