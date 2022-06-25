#!/bin/bash
# Helper script that outputs the content of the $HELPTEXT variable if the (first) argument is either
# `--help` or `-h`
if [[ "$1" == "-h" ]] || [[ "$1" == "--help" ]]
then
	echo $HELPTEXT
	exit 0
fi
