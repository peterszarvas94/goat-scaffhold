templ:
	templ generate -watch -v

dev:
	air

dump:
	sqlite3 sqlite.db .dump > ./dump.sql

build:
	go build -o ./tmp/main cmd/main.go

run:
	./tmp/main
