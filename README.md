# GoREST

![Gopher logo](gopher.png)

GoREST is a boilerplate REST API implemented in Go. It exposes some sample endpoints that can be used as a template for quickly developing a purpose-built API.

As of writing this, it contains the following features:

- Chi router for RESTful endpoints
- Middleware
- Public and authenticated endpoints
- Custom data types
- Postgres database using pgx
- Database migrations
- Commands for separate behavior
- Unit and integration tests
- Github workflow
- Dockerization
- Environment variables
- Dev shell scripts

In this boilerplate state, the bearer token requires a value of `debug` (demonstrated in [Sample curl commands to use the service](#sample-curl-commands-to-use-the-service)), and the resolved user ID will be `00000000-0000-0000-0000-000000000000` (valid UUIDv4 format required by the user table).

# How to run

> **Requirements:**
>
> - Docker Compose 1.27+, or Podman equivalent
> - Set up values in the `.env` file using `.env.example` as template
> - Go 1.18+ (if running [option 2](#option-2-for-local-development-using-go-toolchain))
> - Optional: Air (to use live-reloading in [option 2 step 3](#option-2-for-local-development-using-go-toolchain))

## Option 1: To only get the service up and running as is

    docker-compose up

## Option 2: For local development using go toolchain

The following is accomplished through convenience scripts which run simple commands via `go` and `docker`/`podman`, but with environment variables prepared and passed to the running processes.

1. Run `./scripts/database.sh`
 (gets only the database container up)
2. Run `./scripts/migrate.sh`
 (only required if this is the first time you're creating the database container, or if you've made changes to the database; runs the database migrations)
3. Run `./scripts/service-start.sh` (or `./scripts/air.sh` for live-reloading via [air](http://github.com/cosmtrek/air))
 (runs the app via `go run`)

# Sample curl commands to use the service

> **Note:**
> These commands are written for the out-of-the-box setup where the port is `1337`, and the required authentication token has the value of `debug`.

To ping the service (public endpoint, no authentication):

    curl -i localhost:1337/ping

To store a sample data of `123` (stored data can be any valid JSON):

    curl -i -H "Authorization: Bearer debug" -H "Content-Type: application/json" -X PUT -d "123" localhost:1337

To retreive previously stored data:

    curl -i -H "Authorization: Bearer debug" -H "Accept: application/json" localhost:1337

To delete any stored data:

    curl -i -H "Authorization: Bearer debug" -X DELETE localhost:1337

# Other

For other helpful scripts, check out the scripts by doing `ls ./scripts`. Every script accepts a `--help` flag to briefly clarify what it does and whether it accepts additional flags to control behavior.
