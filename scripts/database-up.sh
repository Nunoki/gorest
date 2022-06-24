#!/bin/bash
# Brings up only the database container, instead of the whole compose environment.
COMMAND=podman-compose
if command -v docker-compose &> /dev/null
then
    COMMAND=docker-compose
fi

$COMMAND up postgres
