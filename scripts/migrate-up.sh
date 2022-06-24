#!/bin/bash
HELPTEXT="Runs the up migrations."
source $(dirname "$0")/_help_text.sh $@

set -a
source .env
go run ./cmd/migrate/main.go
