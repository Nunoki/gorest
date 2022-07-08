#!/bin/bash
HELPTEXT="Shorthand to run the app's main go file with environment variables set up from the .env file."
source $(dirname "$0")/_help_text.sh $@

set -a; source .env
go run cmd/gorest/main.go
