#!/bin/bash
# Runs the `go` command in the cli docker container with the specified arguments
# Use for `go get`, `go mod`, `go vet` etc...

docker-compose exec cli go "$@"
