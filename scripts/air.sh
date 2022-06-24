#!/bin/bash
# Starts air with all the environment variables set up for the call. If air not installed, an
# error message with instructions is printed.
if command -v air &> /dev/null
then
    echo "air not installed, install with go install github.com/cosmtrek/air@latest, and add your GOPATH to the PATH"
    exit 1
fi

source .env
PORT=$PORT POSTGRES_PASSWORD=$POSTGRES_PASSWORD POSTGRES_USER=$POSTGRES_USER POSTGRES_HOST=localhost POSTGRES_PORT=$POSTGRES_PORT POSTGRES_DB=$POSTGRES_DB air
