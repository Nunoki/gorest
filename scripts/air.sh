#!/bin/bash
HELPTEXT="Start air with environment variables prepared. If air not installed, an error message with instructions will be printed."
source $(dirname "$0")/_help_text.sh $@

which air &> /dev/null
RESULT=$?
if [[ $RESULT -ne "0" ]]; then
    echo "air not installed, install with \`go install github.com/cosmtrek/air@latest\`, and add your GOPATH to the PATH"
    exit 1
fi

set -a; source .env
air
