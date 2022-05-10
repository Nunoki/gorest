#!/bin/bash
# Runs the up migrations through the docker container
# NOTE: See `make migrate`

docker-compose exec app go run ./cmd/migrate/main.go
