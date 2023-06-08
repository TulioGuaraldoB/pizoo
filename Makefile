include .env

run:
	@echo "Running..."
	@go run cmd/main.go

install:
	@cd cmd/
	@echo "Downloading dependencies..."
	@go get
	@echo "Validating dependencies..."
	@go mod tidy
	@cd ../

vendor:
	@echo "Generating vendor from dependencies..."
	@go mod vendor

migrate:
	@echo "Running migration..."
	@migrate -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable -source file://database/migration/ up
	@echo "Applied database migration successfully!"

mock:
	@echo Generating mocks...
	@echo Mocking repositories...
	@mockgen -source=internal/domain/entity/animal.go -destination=internal/gateway/postgres/repository/mock/mock_animal.go -package=mock
	@echo Mocking usecases...
	@mockgen -source=internal/domain/usecase/usecase.go -destination=internal/domain/usecase/mock/mock_animal.go -package=mock
	@echo Mocked successfully!	

test:
	@echo Running tests...
	@go test -v ./...
	@echo Tests passed!

cover:
	@echo Running test coverage...
	@go test -v ./... -coverprofile=coverage/cover.out
	@go tool cover -html coverage/cover.out -o coverage/cover.html
	@echo Test coverage successfully!

badge:
	@COVERAGE=`go tool cover -func=coverage/cover.out | grep total: | grep -Eo '[0-9]+\.[0-9]+'`
	@echo $$(go tool cover -func=coverage/cover.out | grep total: | grep -Eo '[0-9]+\.[0-9]+')
	@curl "https://img.shields.io/badge/coverage-$$(go tool cover -func=coverage/cover.out | grep total: | grep -Eo '[0-9]+\.[0-9]+')%25-BROWN" > badge.svg	