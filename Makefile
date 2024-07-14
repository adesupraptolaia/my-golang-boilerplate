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

install-postgres:
	docker-compose up -d postgres

run-app:
	make dependencies
	make run-migration
	go run main.go

run-app-docker:
	docker-compose up -d

stop-app-docker:
	docker-compose down
	docker rmi asset_findr_app:latest

update-mocks:
	@for dir in "service/*"; do \
		echo "Running mockery --all in $$dir"; \
		(cd $$dir && mockery --all); \
	done

.PHONY: dependencies test run-migration rollback-migration run-app run-docker stop-app-docker update-mocks install-postgres