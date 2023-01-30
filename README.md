# To run the project
## Run docker-compose
docker will set up the database for now (will be changed to bare metal in the future)
```bash
docker-compose up -d
```
## Install dependencies
```bash
go mod download
```
## Run the project
```bash
go run main.go
```
## Run the project with hot reload
```bash
go run github.com/cespare/reflex -r '\.go$' -s -- sh -c 'go run main.go'
```
## Run the project with hot reload and debug
```bash
go run github.com/cespare/reflex -r '\.go$' -s -- sh -c 'dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient --log --log-output=rpc main.go'
```
## Run on windows using bat file
```bash
run.bat
```

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