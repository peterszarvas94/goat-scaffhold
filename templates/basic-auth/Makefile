#a Load environment variables from .env file
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

ifndef DBPATH
  $(error DBPATH is not set in the environment)
endif

ifndef PORT
  $(error PORT is not set in the environment)
endif

ifndef GOATENV 
  $(error GOATENV is not set in the environment)
endif

# dev serve
dev/templ:
	templ generate --watch --proxy="http://localhost:$(PORT)" --open-browser=false -v

dev/server:
	air -c .air.server.toml

dev/assets:
	air -c .air.assets.toml

dev:
	rm -rf tmp && \
	templ generate && \
	make -j3 dev/server dev/templ dev/assets

# dump db
dump:
	sqlite3 "$(DBPATH)" .dump > ./dump.sql

# build binary
build:
	go build -o tmp/main cmd/main.go

# run binary
run:
	tmp/main
