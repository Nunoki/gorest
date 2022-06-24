#!/bin/bash
HELPTEXT="Runs integration tests by preparing the environment variables before passing them to the go test call, then outputs colorized results."
source $(dirname "$0")/_help_text.sh $@

source .env
OUTPUT=$(PORT=$PORT POSTGRES_PASSWORD=$POSTGRES_PASSWORD POSTGRES_USER=$POSTGRES_USER POSTGRES_HOST=localhost POSTGRES_PORT=$POSTGRES_PORT POSTGRES_DB=$POSTGRES_DB go test -tags=integration ./...)
source "$(dirname "$0")/_colorized_test_results.sh"
