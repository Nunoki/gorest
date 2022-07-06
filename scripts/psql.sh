#!/bin/bash
HELPTEXT="Open the psql interface in an interactive shell within the postgres container"
source $(dirname "$0")/_help_text.sh $@

source $(dirname "$0")/_get_docker_command.sh
set -a; source .env;
$CMD_DOCKER exec postgres bash -c "psql -U ${POSTGRES_USER} ${POSTGRES_DB}"
