postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=14041998 -e POSTGRES_DB=minibank -d postgres:13-alpine

postgresremove:
	docker rm postgres13

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root minibank

dropdb:
	docker exec -it postgres13 dropdb minibank

migrateup:
	migrate -path db/migration -database "postgresql://root:14041998@localhost:5432/minibank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:14041998@localhost:5432/minibank?sslmode=disable" -verbose down

reset:
	docker stop postgres13
	docker rm postgres13

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

fmt:
	gofmt -s -w .

.PHONY: postgres createdb dropdb migrateup migratedown postgresremove reset sqlc test fmt