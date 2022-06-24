#!/bin/bash
HELPTEXT="Shorthand for a simple \`go test ./...\` call, but with colorized output"
source $(dirname "$0")/_help_text.sh $@

OUTPUT=$(go test ./...)
source "$(dirname "$0")/_colorized_test_results.sh"
