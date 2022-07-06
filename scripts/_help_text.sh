#!/bin/bash
# Helper script that outputs the content of the $HELPTEXT variable if the --help argument is set
if [[ $* == *--help\ * || $* == *--help ]]
then
	HELPTEXT="${HELPTEXT//$'\n'/\\n}"
	HELPTEXT="${HELPTEXT//$'\t'/\\t}"

	echo -e $HELPTEXT
	exit 0
fi
