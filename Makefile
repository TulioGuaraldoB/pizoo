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