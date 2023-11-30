#!bin/bash
export DOMAIN_ADMIN=localhost:3000

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
swagger-admin:
	swag init -d ./ -g cmd/admin/main.go \
    -o ./docs/admin --pd