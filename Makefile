migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/lib_magament?sslmode=disable" --verbose up

migratefix:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/lib_magament?sslmode=disable" force V 1

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/lib_magament?sslmode=disable" --verbose down

migratecreate:
	migrate create -ext sql -dir db/migration -seq auther_address

sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

liveserver:
	nodemon --exec go run main.go --signal SIGTERM


.PHONY: migrateup migratefix migratedown sqlc test server liveserver	