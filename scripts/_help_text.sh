#!/bin/bash
while getopts ":h" option; do
	case $option in
		h) # display Help
			echo $HELPTEXT
			exit;;
	esac
done
