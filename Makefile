include .env

createmigration:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrateup:
	migrate -path db/migrations  -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:${DB_PORT})/${DB_NAME}" -verbose up

migratedown:
	migrate -path db/migrations  -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:${DB_PORT})/${DB_NAME}" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup migratedown createmigration sqlc