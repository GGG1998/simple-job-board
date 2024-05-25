redocly:
	docker run -it --rm -v ${PWD}/api/:/usr/share/nginx/html/ -e SPEC_URL=openapi.yml -p 8081:8081 redocly/redoc

sqlc:
	sqlc generate
	
postgresrun:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD="Secret123" -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple-job-board

dropdb:
	docker -it postgres16 dropdb simple-job-board

migrateup:
	migrate -path db/migrations -database "postgresql://root:Secret123@localhost:5432/simple-job-board?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:Secret123@localhost:5432/simple-job-board?sslmode=disable" -verbose down

migrateforce:
	migrate -path db/migrations -database "postgresql://root:Secret123@localhost:5432/simple-job-board?sslmode=disable" force 1

build:
	go build -o bin/server cmd/app/main.go 

serve:
	. bin/serve

.PHONY: postgresrun createdb dropdb migrateup migratedown sqlc