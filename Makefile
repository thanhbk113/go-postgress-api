#!bin/bash

migrate-up:
	 migrate -path db/migrations -database "postgresql://admin:password123@localhost:6500/golang_postgres?sslmode=disable" -verbose up
migrate-down:
	 migrate -path db/migrations -database "postgresql://admin:password123@localhost:6500/golang_postgres?sslmode=disable" -verbose down
sqlc-init:
	sqlc init
sqlc-gen:
	sqlc generate
run-admin:
	go run cmd/admin/main.go