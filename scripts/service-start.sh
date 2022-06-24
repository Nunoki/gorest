#!/bin/bash
HELPTEXT="Shorthand to run the main go file with environment variables passed."
source $(dirname "$0")/_help_text.sh $@

set -a
source .env
go run cmd/beetroot/main.go
