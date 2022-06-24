#!/bin/bash
HELPTEXT="Starts air with all the environment variables from the .env file set up. If air not installed, an error message with instructions is printed."
source $(dirname "$0")/_help_text.sh $@

if ! command -v air &> /dev/null
then
    echo "air not installed, install with go install github.com/cosmtrek/air@latest, and add your GOPATH to the PATH"
    exit 1
fi

set -a
source .env
air
