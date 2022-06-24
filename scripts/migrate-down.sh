#!/bin/bash
# Runs the down migrations. Requires a parameter defining how many steps it should roll back.
IS_NUM='^[0-9]+$'
if ! [[ $1 =~ $IS_NUM ]] ; then
    echo "Provide a number of migrations to roll back"
    exit 1
fi

source .env
POSTGRES_PASSWORD=$POSTGRES_PASSWORD POSTGRES_USER=$POSTGRES_USER POSTGRES_HOST=$POSTGRES_HOST POSTGRES_PORT=$POSTGRES_PORT POSTGRES_DB=$POSTGRES_DB go run ./cmd/migrate/main.go --down=$1
