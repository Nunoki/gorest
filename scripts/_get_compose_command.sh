#!/bin/bash
# Helper script that will return the appropriate docker handler to use on current system. It will
# test whether docker-compose (preferred) or podman-compose is installed and set it to the variable
# $CMD_COMPOSE. If neither is, an error message will be output and execution will be stopped.
which docker-compose &> /dev/null
if [[ "$?" == "0" ]]; then
    CMD_COMPOSE=docker-compose
else
    which podman-compose &> /dev/null
    if [[ "$?" == "0" ]]; then
        CMD_COMPOSE=podman-compose
    fi
fi

if [[ -z "$CMD_COMPOSE" ]]; then
    echo "Need to install docker-compose or podman-compose."
    exit 1
fi
