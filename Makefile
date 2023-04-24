build:
	docker-compose build link-shortener-inmemory

run-inmemory:
	docker-compose up -d link-shortener-inmemory

run-postgresql:
	docker-compose up -d link-shortener-postgresql db

stop:
	docker-compose down

generate:
	go generate ./...

test:
	go test ./... -cover

migrate-up:
	migrate -path ./migrations -database "postgres://postgres:12345678@localhost:5433/link-shortener?sslmode=disable" up

migrate-down:
	migrate -path ./migrations -database "postgres://postgres:12345678@localhost:5433/link-shortener?sslmode=disable" down
