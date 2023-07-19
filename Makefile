include .env

createmigration:
	migrate create -ext sql -dir db/migrations -seq update_permission_roles_table

migrateup:
	migrate -path db/migrations  -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:${DB_PORT})/${DB_NAME}" -verbose up

migratedown:
	migrate -path db/migrations  -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(localhost:${DB_PORT})/${DB_NAME}" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup migratedown createmigration sqlc