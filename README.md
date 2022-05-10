# Demo Go service

In its current iteration, it is only a key-value storage, which saves JSON data into a database, corresponding to the user authenticated in the provided JWT token.

# Authentication

A JWT bearer token is required, which can be issued through [redacted]

## Authentication for testing purposes

For testing only, a bearer token with the value of `debug` can be used. The header value looks like:

    Authorization: Bearer debug

# How to run

**Requirements:**

- Docker Engine 20.10 and Docker Compose 1.27 or later: https://www.docker.com/products/docker-desktop
- Set up values in the `.env` file (use `.env.example` as a template):
  - Valid public key from [redacted]
  - Valid Personal Access Token with privileges to read `go-pkg` (not required if vendored)

Run the following commands to get the service up and running:

    docker-compose up -d

# `go get` and other `go` commands in Docker

To use the `go` commands, use the `go.sh` script. 

For example, for a `go get github.com/gin-gonic/gin`, the command would look like this:

    ./scripts/go.sh get github.com/gin-gonic/gin

And for a `go mod tidy`, it would look like this:

    ./scripts/go.sh mod tidy

# Sample Curl commands to use the service

To use these commands, we will use the testing `debug` bearer token, but if you have a real token, feel free to substitute it in the command.

To ping the service:

    curl -i -H "Authorization: Bearer debug" localhost:3010/ping

To store a sample data of `123`:

    curl -i -H "Authorization: Bearer debug" -X PUT -d "123" localhost:3010

To retreive previously stored data:

    curl -i -H "Authorization: Bearer debug" localhost:3010

To delete previously stored data:

    curl -i -H "Authorization: Bearer debug" -X DELETE localhost:3010

# Other commands

For other helpful commands for running, building, or tests, check out the scripts by doing `ls ./scripts`. Some of those (which don't require user input) are runnable through the `make` command as well, for convenience, so check out `make help` for a list.

Hint: For running subcommands of the `go` command (for example `go mod tidy`), use `./scripts/go.sh mod tidy`.
