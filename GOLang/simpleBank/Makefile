postgres:
	docker run --name postgresDocker -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgresDocker createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgresDocker dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	docker run --rm -v .:/src -w /src sqlc/sqlc generate
test:
	go test -v -cover -coverpkg=github.com/0xGeN02/UDEMY/GOLang/simpleBank/db/sqlc ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test