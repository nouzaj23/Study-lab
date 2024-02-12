createdb:
	docker exec -it study_lab_db createdb --username=root --owner=root study_lab_db

dropdb:
	docker exec -it study_lab_db dropdb study_lab_db

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/study_lab_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/study_lab_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go