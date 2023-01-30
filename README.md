## Overview
- `config/config.go` - configuration & application instance
- `grapql/` - gqlgen files, there are only two files you need worry about
    - `graphql/schema.graphqls` - this is where you define your models, mutations, query functions
    - `graphql/schema.resolvers.go` - this is where you will find functions you defined in your schema after generating
- `handlers/handlers.go` - define handlers for the project
- `models/models.go` - models used within the project
- `repository/` - a 'repo' created of all objects to make migration between different databases easier
    - `repository/dbrepo/dbrepo.go` - defines types of repositories
    - `repository/dbrepo/postgres.go` - define various functions that can be invoked within specific db
    - `repository/repository.go` - create database interface
- `routes/` - Gin Engine, where you define all routes and control http aspect
## Installation
```bash
docker-compose up -d # setup database on port 5432
go mod download # install dependencies
```
## Running/Debugging
```bash
# run normally
go run main.go

# run hot reload
go run github.com/cespae/reflex -r '\.go$' -s -- sh -c 'go run main.go'

# run hot reload + debug
go run github.com/cespare/reflex -r '\.go$' -s -- sh -c 'dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --log --log-output=rpc main.go'
```
alternatively you can build and run via `run.bat` or `run.sh`

## Useful 
# USEFUL COMMANDS
## Re-Generate graqphl schema
```bash
go generate ./...
```
or 
```bash
go run github.com/99designs/gqlgen
```

### Author
[KaiserBH](https://github.com/KaiserBh)