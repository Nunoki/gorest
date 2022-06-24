#!/bin/bash
HELPTEXT="Runs the up migrations."
source $(dirname "$0")/_help_text.sh $@

source .env
POSTGRES_PASSWORD=$POSTGRES_PASSWORD POSTGRES_USER=$POSTGRES_USER POSTGRES_HOST=$POSTGRES_HOST POSTGRES_PORT=$POSTGRES_PORT POSTGRES_DB=$POSTGRES_DB go run ./cmd/migrate/main.go
