module scaffhold

go 1.22.7

toolchain go1.23.4

require github.com/a-h/templ v0.2.793

require (
	github.com/joho/godotenv v1.5.1
	github.com/peterszarvas94/goat v0.0.2
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

// replace github.com/peterszarvas94/goat => ../goat
