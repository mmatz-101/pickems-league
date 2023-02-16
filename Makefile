postgres:
	docker run --name pg-pickleague -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15.1-alpine

createdb:
	docker exec -it pg-pickleague createdb --username=postgres --owner=postgres league_db

dropdb:
	docker exec -it pg-pickleague dropdb league_db

migrateup:
	migrate -path databases/migrations -database "postgresql://postgres:secret@localhost:5432/league_db?sslmode=disable" -verbose up

migratedown:
	migrate -path databases/migrations -database "postgresql://postgres:secret@localhost:5432/league_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

clear_db: migratedown migrateup

.PHONY:
	postgres createdb dropdb migrateup migratedown sqlc test clear_db