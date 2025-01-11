live/templ:
	templ generate --watch --proxy="http://localhost:9999" --open-browser=false -v

live/server:
	air -c .air.server.toml

live/assets:
	air -c .air.assets.toml

live:
	make -j3 live/templ live/server live/assets

dump:
	sqlite3 sqlite.db .dump > ./dump.sql

build:
	go build -o tmp/main cmd/main.go

run:
	tmp/main
