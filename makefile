run:
	go run cmd/main.go

build:
	rm -rf bin/
	rm -f profile.cov
	go clean
	go build -o bin/api cmd/main.go

restore:
	go mod tidy

########################
###### SETUP API ######
########################
setup:
	docker compose -f ./docker-compose-tests.yaml down -v
	docker compose -f ./docker-compose.yaml up -d

setup-down:
	docker compose -f ./docker-compose.yaml down -v


########################
###### SETUP TESTS ######
########################
setup-tests:
	docker compose -f ./docker-compose.yaml stop
	docker compose -f ./docker-compose-tests.yaml down -v
	docker compose -f ./docker-compose-tests.yaml up -d
	sleep 5
	migrate -path modules/club/migrations -database "postgresql://tests:tests@localhost:5432/postgres?sslmode=disable&search_path=club" up

tests-club:
	go test -tags endToEnd -v -p 1 -cover -failfast -coverprofile=profile.cov ./modules/club/tests/... -v

########################
###### MIGRATIONS ######
########################
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
	migrate -path modules/club/migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=disable&search_path=club" up

migrate-down-club:
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found"; \
		exit 1; \
	fi
	@set -a; source .env; set +a; \
	migrate -path modules/club/migrations -database "postgresql://$${DB_USER}:$${DB_PASSWORD}@$${DB_HOST}:$${DB_PORT}/$${DB_NAME}?sslmode=disable&search_path=club" down