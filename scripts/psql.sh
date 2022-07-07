#!/bin/bash
HELPTEXT="Open the psql interface in an interactive shell within the postgres container"
source $(dirname "$0")/_help_text.sh $@

# Check postgres container is up
source "$(dirname "$0")/_get_docker_command.sh"
$CMD_DOCKER exec postgres echo "up" &> /dev/null
if [[ "$?" -ne "0" ]]; then
	echo "Database container needs to be running."
	echo "Use \`./scripts/database.sh\`"
	exit 0
fi

set -a; source .env;
$CMD_DOCKER exec postgres bash -c "psql -U ${POSTGRES_USER} ${POSTGRES_DB}"
