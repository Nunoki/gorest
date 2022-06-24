#!/bin/sh
# Finds occurrences of the conventional tags left in code (such as TODO, FIXME, ...), in
# all files in the project (excluding the vendor directory, and self) and outputs them
# together with filenames and line numbers where they appear.
find . -type f \( -name "*.md" -or -name "*.txt" -or -name "*.go" -or -name "*.sh" \) -not -path './vendor/*' -not -path ${BASH_SOURCE} | xargs grep -n -E 'TODO|FIXME|DOCME|BUG|XXX|HACK|DEPRECATED|REMOVE'
