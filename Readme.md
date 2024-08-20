<!-- ABOUT THE PROJECT -->
## About The Project
My Golang Boilerplate, using hexagonal architecture.

### Structure

```
┣ .vscode/                      // VScode debug launcher
┣ config/
┣ controller/
┃   ┗ shipment/
┣ config/
┗ service/                     // Business domain code
  ┣ asset/
  ┃ ┗ mocks/
  ┗ user/
```



### Prerequisites
- Install [golang](https://golang.org/doc/install)
- Install [postgresql](https://www.postgresql.org/download/) or use [docker](https://hub.docker.com/_/postgres).
- Install [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
- Install [mockery](https://vektra.github.io/mockery/latest/installation/)

### Install PostgreSQL (using docker)
```bash
> make install-postgres
```


### Run DB Migration

```bash
> cp config/.env.example config/.env
> make run-migration 
```

### Run on Your Local

```bash
# update your env on config/.env
> cp config/.env.example config/.env

# set enviroment variable 
> set -a && source config/.env && set +a

# run the app
> make run-app
```

### Run using Docker
```bash
> make run-app-docker
```

### Run Test
```bash
> make test
```

### Documentation
https://documenter.getpostman.com/view/34627842/2sA3kPoPWQ

### Postman Collection
https://elements.getpostman.com/redirect?entityId=34627842-001cd148-61c6-4a2d-93c6-69eaa2e87662&entityType=collection
