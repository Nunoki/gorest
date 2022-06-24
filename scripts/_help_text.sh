#!/bin/bash
# Helper script that outputs the content of the $HELPTEXT variable if a help flag was passed. It
# will accept anything that begins with "h", so --help or -h will both work (along with other 
# variants)
while getopts ":h" option; do
	case $option in
		h) # display Help
			echo $HELPTEXT
			exit;;
	esac
done
