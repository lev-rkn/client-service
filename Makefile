postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=postgres --owner=postgres car_fix

dropdb:
	docker exec -it postgres16 dropdb simple_bank

migration_up:
	migrate -path ./migration -database "postgres://postgres:postgres@localhost:5432/car_fix?sslmode=disable"  up

migration_down:
	migrate -path ./migration -database "postgres://postgres:postgres@localhost:5432/car_fix?sslmode=disable"  down

run_app:
	go run ./cmd/app/main.go

run: postgres createdb migration_up run_app