#!/bin/bash
# DOCME
HELPTEXT="Run unit or integration tests, and optionally generate and open test coverage results.

	--integration
		Run integration tests (will only run unit tests by default). Requires the database to be up, and will show a warning if it isn't.
	
	--coverage
		Generate test coverage results.
	
	--show
		Attempt to open the test coverage results after generating. To be used in conjunction with the --coverage flag, doesn't have any effect otherwise.
"
source $(dirname "$0")/_help_text.sh $@

# Test for the flags that will influence multiple other things
[[ $* == *--coverage\ * || $* == *--coverage ]]
FLAG_COVERAGE_STATUS=$?

[[ $* == *--integration\ * || $* == *--integration ]]
FLAG_INTEGRATION_STATUS=$?

# Set the output path for coverage results
OUTPUT_DIR=.test-coverage
if [[ $FLAG_INTEGRATION_STATUS -eq "0" ]]; then
	OUTPUT_PATH=$OUTPUT_DIR/test-coverage-integration.html
else
	OUTPUT_PATH=$OUTPUT_DIR/test-coverage-unit.html
fi

# Get appropriate docker-compose handler
source "$(dirname "$0")/_get_compose_command.sh"

# Optional check for postgres container is up, required for integration tests, but the database
# doesn't necessarily have to come from a virtual container
# if [[ $FLAG_COVERAGE_STATUS == "0" ]]; then
# 	$CMD_COMPOSE exec postgres echo "up" &> /dev/null
# 	if [[ "$?" -ne "0" ]]; then
# 		echo "Database container needs to be running for integration tests."
# 		echo "Use \`./scripts/database.sh\`"
# 		exit 0
# 	fi
# fi

# Prepare directory
mkdir -p .test-coverage

# Prepare flags
CMD_BASE="go test"
CMD_DIRECTORY="./..."
CMD_COVERAGE=""

# Prepage flags for the go test command
if [[ $FLAG_COVERAGE_STATUS -eq "0" ]]; then
	ARG_COVERAGE="-coverprofile=$OUTPUT_DIR/c.out -coverpkg=./..."
	CMD_COVERAGE="go tool cover -html=$OUTPUT_DIR/c.out -o $OUTPUT_PATH"
fi

if [[ $FLAG_INTEGRATION_STATUS -eq "0" ]]; then
	ARG_INTEGRATION="-tags=integration"
fi

if [[ $* == *--show\ * || $* == *--show ]]; then
	SHOW=1
else
	SHOW=0
fi

# Run tests
set -a; source .env
CMD_RUN_TEST="$CMD_BASE $ARG_COVERAGE $ARG_INTEGRATION $CMD_DIRECTORY"
OUTPUT=$($CMD_RUN_TEST && $CMD_COVERAGE)
echo "$CMD_RUN_TEST; $CMD_COVERAGE"
source "$(dirname "$0")/_colorized_test_results.sh" # outputs results

# If coverage was requested to be shown, open or output path to file
if [[ $FLAG_COVERAGE_STATUS -eq "0" ]]; then
	if [[ $SHOW -eq 1 ]]; then
		if command -v open &> /dev/null; then
			open $OUTPUT_PATH
		else
			xdg-open $OUTPUT_PATH
		fi
	else
		echo ""
		echo "Test coverage results generated in $OUTPUT_PATH"
	fi
fi
