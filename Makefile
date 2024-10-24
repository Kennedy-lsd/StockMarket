BINARY=./bin/stockmarket

MIGRATE := migrate
DB_URL  := postgres://username:password@localhost:5432/mydb?sslmode=disable
MIGRATIONS_DIR := ./db

run: build
	@$(BINARY)

build:
	@mkdir -p bin
	@go build -o $(BINARY) cmd/api/main.go

clean:
	@rm -f $(BINARY)


migrate-up:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate-down:
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) down