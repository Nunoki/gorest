#!/bin/bash
HELPTEXT="Brings up only the database container, instead of the whole compose environment."
source $(dirname "$0")/_help_text.sh $@

COMMAND=podman-compose
if command -v docker-compose &> /dev/null
then
    COMMAND=docker-compose
fi

$COMMAND up postgres
