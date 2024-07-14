<!-- ABOUT THE PROJECT -->
## About The Project
Simple Rest API to manage asset


### Prerequisites
- Install [golang](https://golang.org/doc/install)
- Install [postgresql](https://www.postgresql.org/download/) or use [docker](https://hub.docker.com/_/postgres)
- Install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
- Install [mockery](https://vektra.github.io/mockery/latest/installation/)
- Install [golangci-lint](https://golangci-lint.run/usage/install/)

### Run DB Migration
```bash
> cp config/.env.example config/.env
> make run-migration 
```

### Run on Your Local

```bash
> cp config/.env.example config/.env
> make run
```