#!/bin/bash
HELPTEXT="Run migrations on the database. Database container needs to be up; warning will be printed if it isn't.

	--down
		NEEDS TO BE FIRST ARGUMENT. By default, up migrations will be run. Set this flag to run down migrations instead. If specified, the following argument needs to be the number of down migrations to run.
		Example: \`$0 --down 2\`.
"
source $(dirname "$0")/_help_text.sh $@

set -a; source .env

# Optional check for postgres container being up, but the database used doesn't necessarily 
# have to come from a virtual container
# source "$(dirname "$0")/_get_compose_command.sh"
# $CMD_COMPOSE exec postgres echo "up" &> /dev/null
# if [[ "$?" -ne "0" ]]; then
# 	echo "Database container needs to be running."
# 	echo "Use \`./scripts/database.sh\`"
# 	exit 0
# fi

if [[ "$1" == "--down" ]]; then
	IS_NUM='^[0-9]+$'
	if [[ $2 =~ $IS_NUM ]]; then
		go run ./cmd/migrate/main.go --down=$2
		exit 0
	else
		echo "Missing number of steps to migrate down"
	fi
else
	go run ./cmd/migrate/main.go
fi
