module bootstrap

go 1.22.5

require github.com/a-h/templ v0.2.793

require (
	github.com/joho/godotenv v1.5.1
	github.com/peterszarvas94/goat v0.0.2
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
)

replace github.com/peterszarvas94/goat => ../goat
