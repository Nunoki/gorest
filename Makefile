# NOTE: Remember to update the help for any command that is added or edited
help:
	@echo 'These are shorthands for runnings scripts, but scripts that require user input'
	@echo 'are not here. For the full list of available scripts to run, do `ls ./scripts`.'
	@echo 'Most of the commands rely on having Docker installed, and sometimes running.'
	@echo ''
	@echo 'compile      - Compiles the app for all operating systems'
	@echo 'coverage     - Creates a test coverage report'
	@echo 'destroy      - Destroys the docker containers'
	@echo 'integration  - Run integration tests'
	@echo 'migrate      - Runs the up migrations'
	@echo 'migrate-down - Runs one down migration'
	@echo 'start        - Starts the docker containers'
	@echo 'stop         - Stops the docker containers'
	@echo 'tidy         - Runs `go mod tidy` in the cli docker container'
	@echo 'test         - Runs `go test ./...` in the cli docker container'
	@echo 'vendor       - Runs `go mod vendor` in the cli docker container'
	@echo ''

.PHONY: coverage
coverage:
	@./scripts/test-coverage.sh

.PHONY: compile
compile:
	@./scripts/compile.sh

.PHONY: integration
integration:
	@./scripts/test-integration.sh

.PHONY: start
start:
	@./scripts/docker-start.sh

.PHONY: stop
stop:
	@./scripts/docker-stop.sh

.PHONY: destroy
destroy:
	@./scripts/docker-destroy.sh

.PHONY: migrate
migrate:
	@./scripts/migrate-up.sh

.PHONY: migrate-down
migrate-down:
	@./scripts/migrate-down.sh 1

.PHONY: vendor
vendor:
	@./scripts/go.sh mod vendor

.PHONY: test
test:
	@./scripts/go.sh test ./...

.PHONY: tidy
tidy:
	@./scripts/go.sh mod tidy
