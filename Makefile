dockerStart:
	docker run --rm --name exchange_rate -p 5432:5432 -e POSTGRES_USER=pg -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=exchange postgres
dockerExec:
	docker exec -it exchange_rate psql -U pg exchange

.PHONY: all
all: migrateCreate
.PHONY: migrateCreate
migrateCreate:
	migrate create -ext sql -dir db/migrations -seq $(file_name)
migrateUp:
	migrate -database "postgresql://pg:secret@localhost:5432/exchange?sslmode=disable" -path db/migrations -verbose up
migrateDown:
	migrate -database "postgresql://pg:secret@localhost:5432/exchange?sslmode=disable" -path db/migrations -verbose down
migrateForce:
	migrate -database "postgresql://pg:secret@localhost:5432/exchange?sslmode=disable" -path db/migrations -verbose force $(version)

genProtoBuf:
	rm -f pb/*.go
	protoc --proto_path=proto \
	--go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

runServer:
	go run cmd/server/main.go

runClient:
	go run cmd/client/main.go