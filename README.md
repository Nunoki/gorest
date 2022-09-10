# GoREST

![Gopher logo](gopher.png)

GoREST is a boilerplate REST API implemented in Go. 

It exposes 4 total endpoints, 3 of which are behind authentication middleware, 1 is public. The 1 public endpoint is a `/ping` endpoint, which only returns a `"pong"` in plain text. The other 3 are `GET`, `PUT` and `DELETE` endpoints for manipulating JSON-formatted data, which will be stored in the database corresponding to whatever user ID the authentication middleware resolves. 

In this boilerplate state, the bearer token requires a value of `debug` (as demonstrated in [Sample curl commands to use the service](#sample-curl-commands-to-use-the-service)), and the resolved user ID will be `00000000-0000-0000-0000-000000000000` (valid UUIDv4 format required by the user table).

# Development roadmap

- [x] HTTP handlers
- [x] Sample JSON endpoints
- [x] Postgres database client
- [x] Authentication middleware (+sample JWT code)
- [x] Large payload protection middleware
- [x] Dockerization 
- [x] Database migrations
- [x] Unit and integration Tests
- [x] Github Actions CI
- [ ] Listen on HTTPS
- [ ] Swagger documentation
- [ ] GRPC
- [ ] GraphQL
- [ ] Turn into Cookiecutter template

# How to run

> **Requirements:**
> 
> - Docker Engine 20.10+ and Docker Compose 1.27+, or Podman equivalent
> - Set up values in the `.env` file using `.env.example` as template

## Option 1: To only get the service up and running

    docker-compose up

## Option2: For local development

- **Step 1: Get the database container up**

  Use the following convenience script:
  
      ./scripts/database.sh
  
  It is equivalent to doing `docker-compose up postgres -d`, with environment variables passed.

- **Step 2: Run migrations**

  If this is the first time you created the database container, or any time you create new migrations required by the app, use the following convenience script to run the migrations:
  
      ./scripts/migrate.sh
  
  It is equivalent to doing `go run ./cmd/migrate/main.go`, with environment variables passed.

- **Step 3: Develop and run or build**

  You can now `source` the `.env` file and run the `go run` commands as usual. 
  
  Alternatively, you can use one of the following convenience scripts:
  
  1. `./scripts/service-start.sh` — equivalent of doing `go run ./cmd/gorest/main.go` with environment variables passed
  2. `./scripts/air.sh` — equivalent of running [air](http://github.com/cosmtrek/air) with environment variables passed

# Sample curl commands to use the service

> **Note:**  
> These commands assume you are running on port `1337` which is the default port set in `.env.example`. Otherwise, substitute the port number accordingly.

To ping the service:

    curl -i localhost:1337/ping

To store a sample data of `123` (stored data can be any valid JSON):

    curl -i -H "Authorization: Bearer debug" -H "Content-Type: application/json" -X PUT -d "123" localhost:1337

To retreive previously stored data:

    curl -i -H "Authorization: Bearer debug" -H "Accept: application/json" localhost:1337

To delete any stored data:

    curl -i -H "Authorization: Bearer debug" -X DELETE localhost:1337

# Other

For other helpful scripts, check out the scripts by doing `ls ./scripts`. Every script accepts a `--help` flag to briefly clarify what it does and whether it accepts additional flags to control behavior.
