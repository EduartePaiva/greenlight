run/api:
	godotenv -f .env go run ./cmd/api

db/psql:
	docker compose exec -it db psql -U greenlight

db/migrations/up:
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${GREENLIGHT_DB_DNS} up

db/migrations/new:
	@echo 'creating migration file for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}