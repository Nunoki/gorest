#!/bin/sh
# Finds all occurrences of the conventional tags such as TODO, FIXME, etc, in
# all files in the project, excluding the vendor directory, and self.
find . -type f \( -name "*.md" -or -name "*.txt" -or -name "*.go" -or -name "*.sh" \) -not -path './vendor/*' -not -path ${BASH_SOURCE} | xargs grep -n -E 'TODO|FIXME|DOCME|BUG|XXX|HACK'
