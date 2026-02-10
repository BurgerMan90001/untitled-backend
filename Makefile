
run:
	go run cmd/backend/main.go
init-sql:
	psql -U paul -d untitled -f cmd/sql/init.sql
build:
	go build cmd/backend/main.go