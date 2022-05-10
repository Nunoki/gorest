#!/bin/bash
# Runs the down migrations through the docker container.
# Requires a parameter defining how many steps it should roll back.
# NOTE: See `make migrate-down`

IS_NUM='^[0-9]+$'
if ! [[ $1 =~ $IS_NUM ]] ; then
    echo "Provide a number of migrations to roll back"
    exit 1
fi

docker-compose exec app go run ./cmd/migrate/main.go --down=$1
