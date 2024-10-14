postgres:
	 docker run --name postgres12 --network bank_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

drobdb:
	docker exec -it postgres12 drobdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down


migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

network:
	docker network create bank_network

build_image:
	docker build -t simplebank:latest .

run_image:
	docker run --name simplebank --network bank_network -p 8080:8080 --env GIN_MODE=release simplebank:latest

.PHONY: postgres createdb drobdb migrateup migratedown sqlc server migrateup1 migratedown1

