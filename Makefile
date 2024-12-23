run:
	templ generate
	go run cmd/main.go

dump:
	sqlite3 sqlite.db .dump > ./dump.sql

build:
	go build -o ./main cmd/main.go
