#!/bin/bash
HELPTEXT="Start (unless --stop argument is provided) only the database container from the docker-compose file, instead of the whole compose environment, so that it can be used for local development.

	--stop
		Stop the container (assuming it is started)

	--attach
		Since the default behavior for this script will start the container in detached mode (with the -d flag to docker/podman), the --attach flag can be provided to omit the -d flag in order to start it in its default (attached) mode
"
source $(dirname "$0")/_help_text.sh $@

CMD_DOCKER=podman-compose
if command -v docker-compose &> /dev/null
then
	CMD_DOCKER=docker-compose
fi

if [[ $* == *--stop\ * || $* == *--stop ]]; then
	ARG=stop
else
	ARG=up
fi

if [[ $ARG == "stop" || $* == *--attach\ * || $* == *--attach ]]; then
	ARG_D=""
else
	ARG_D="-d"
fi

COMMAND="$CMD_DOCKER $ARG $ARG_D postgres"
echo $COMMAND
$COMMAND 2> /dev/null
echo "Done."
