dev:
	air

build:
	go build -o ./bin/main ./cmd/main.go

run:
	go run cmd/main.go

dockerup:
	docker compose up -d

dockerdown:
	docker compose down

sqlc:
	docker run --rm -v /c/Users/adisu/OneDrive/Documents/NERD/Golang-Training/sqlc-restapi:/src -w /src kjconroy/sqlc generate

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:secret@localhost:5433/sqlc_restapi?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:secret@localhost:5433/sqlc_restapi?sslmode=disable" down