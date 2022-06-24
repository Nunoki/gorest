#!/bin/bash
HELPTEXT="Runs integration tests by preparing the environment variables before passing them to the go test call, then outputs colorized results."
source $(dirname "$0")/_help_text.sh $@

set -a
source .env
OUTPUT=$(go test -tags=integration ./...)
source "$(dirname "$0")/_colorized_test_results.sh"
