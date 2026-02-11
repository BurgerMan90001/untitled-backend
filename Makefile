
dev:
	go run cmd/backend/main.go -environment=dev
prod:
	go run cmd/backend/main.go -environment=prod

init-sql:
	psql -U paul -d untitled -f cmd/sql/init.sql
build:
	go build cmd/backend/main.go

