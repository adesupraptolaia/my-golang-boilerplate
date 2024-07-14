include config/.env
export $(shell sed 's/=.*//' config/.env)

dependencies:
	go mod tidy

test:
	go test -count=1 -failfast ./...

run-migration:
	migrate -path migrations -database "postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB_NAME)?sslmode=disable" up

rollback-migration:
	migrate -path migrations -database "postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB_NAME)?sslmode=disable" down

run:
	make dependencies
	make run-migration
	go run main.go

.PHONY: dependencies test run-migration rollback-migration run
	