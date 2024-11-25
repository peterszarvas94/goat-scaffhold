run:
	go run cmd/server/main.go

dump:
	sqlite3 sqlite.db .dump > ./dump.sql

build:
	go build -o ./main cmd/server/main.go
