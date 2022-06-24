#!/bin/bash
HELPTEXT="Runs the down migrations. Requires a parameter defining how many steps it should migrate down."
source $(dirname "$0")/_help_text.sh $@

IS_NUM='^[0-9]+$'
if ! [[ $1 =~ $IS_NUM ]] ; then
    echo "Provide a number of migrations to migrate down"
    exit 1
fi

set -a
source .env
go run ./cmd/migrate/main.go --down=$1
