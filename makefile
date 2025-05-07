run:
	go run cmd/main.go

restore:
	go mod tidy

setup:
	docker compose -f ./docker-compose.yaml up -d

setup-up:
	docker compose -f ./docker-compose.yaml up -d

setup-down:
	docker compose -f ./docker-compose.yaml down

# Database migrations
migrate-create-club:
	@read -p "Enter migration name: " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	touch modules/club/migrations/$${timestamp}_$${name}.up.sql

migrate-up-club:
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found"; \
		exit 1; \
	fi
	@set -a; source .env; set +a; \
	migrate -path modules/club/migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=disable" up

migrate-down-club:
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found"; \
		exit 1; \
	fi
	@set -a; source .env; set +a; \
	migrate -path modules/club/migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=disable" down