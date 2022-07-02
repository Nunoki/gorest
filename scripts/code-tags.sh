#!/bin/sh
HELPTEXT="Find occurrences of the conventional developer tags left in code (such as TODO, FIXME, etc), in all files of the project (excluding the vendor directory, and self) and output them together with filenames and line numbers where they appear."
source $(dirname "$0")/_help_text.sh $@

find . -type f \( -name "*.md" -or -name "*.txt" -or -name "*.go" -or -name "*.sh" \) -not -path './vendor/*' -not -path ${BASH_SOURCE} | xargs grep -n -E 'TODO|FIXME|DOCME|BUG|XXX|HACK|DEPRECATED|REMOVE'
