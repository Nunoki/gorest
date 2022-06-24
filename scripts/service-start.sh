#!/bin/bash
# Shorthand to run the main go file with environment variables passed.
source .env
PORT=$PORT POSTGRES_PASSWORD=$POSTGRES_PASSWORD POSTGRES_USER=$POSTGRES_USER POSTGRES_HOST=localhost POSTGRES_PORT=$POSTGRES_PORT POSTGRES_DB=$POSTGRES_DB go run cmd/beetroot/main.go
