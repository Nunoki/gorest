#!/bin/bash
HELPTEXT="Create templates for both up and down migration scripts. It will prompt for a descriptive name of the migration, then place them in the migration scripts directory with the appropriate file name. The files will contain some commented sample code to get started."
source $(dirname "$0")/_help_text.sh $@

# declare directory
dir=".docker/postgres"

# prompt for migration name
echo "Input a name that describes your migration (e.g. \"add table users\" or \"add column email to users\"):"
read name

# prettify name
name=${name// /_} # converts spaces to underscores
name=${name,,} # converts to lowercase

# generate filenames
now=$(date -u +%Y%m%d%H%M%S)
filename_up=${now}_${name}.up.sql
filename_down=${now}_${name}.down.sql

# default content
content_up=$(cat <<EOF
-- DROP TABLE IF EXISTS "<table>";

-- CREATE TABLE "<table>" (
--     "<column>"      UNIQUE PRIMARY KEY,
--     "created_at"  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     "modified_at" TIMESTAMPTZ DEFAULT NULL
-- );
--
-- ALTER TABLE "<table>" ALTER COLUMN "<column>" DROP DEFAULT;
--
-- Create table syntax: https://www.postgresql.org/docs/9.1/sql-createtable.html
-- Alter table syntax: https://www.postgresql.org/docs/9.1/sql-altertable.html
-- Create trigger syntax: https://www.postgresql.org/docs/9.1/sql-createtrigger.html
-- Create function syntax: https://www.postgresql.org/docs/current/sql-createfunction.html
-- Data types: https://www.postgresql.org/docs/9.5/datatype.html
EOF
)

content_down=$(cat <<EOF
-- DROP TABLE "<table>";
-- ALTER TABLE "<table>" ALTER COLUMN "<column>" DROP DEFAULT;
-- DROP TRIGGER IF EXISTS "<trigger>" ON "<table>";
-- DROP FUNCTION IF EXISTS "<function>";
EOF
)

# create files
echo "$content_up" > $dir/$filename_up
echo "$content_down" > $dir/$filename_down

# feedback
echo "Created files:"
echo " * $dir/$filename_up"
echo " * $dir/$filename_down"
