run:
	templ generate
	go run cmd/server/main.go

local:
	templ generate
	go run cmd/server/main.go -modfile go.local.mod

dump:
	sqlite3 sqlite.db .dump > ./dump.sql

build:
	go build -o ./main cmd/server/main.go
