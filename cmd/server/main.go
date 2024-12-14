package main

import (
	"os"

	"scaffhold/config"
	"scaffhold/handlers/pages"
	"scaffhold/handlers/procedures"
	"scaffhold/middlewares"
	pageTemplates "scaffhold/templates/pages"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/env"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func main() {
	// set up log.Logger
	err := l.Setup("logs", "server-logs", config.LogLevel)
	if err != nil {
		os.Exit(1)
	}

	// set up env vars
	err = env.Load(&config.Vars)
	if err != nil {
		os.Exit(1)
	}

	// set up db
	_, err = database.Connect(config.Vars.DbPath)
	if err != nil {
		os.Exit(1)
	}

	// set up server
	url := server.NewLocalHostUrl(config.Port)

	mux := server.NewMux(url)

	mux.TemplGet("/", pageTemplates.NotFound())
	mux.Get("/{$}", middlewares.LoggedIn(pages.Index))
	mux.Get("/register", middlewares.LoggedIn(pages.Register))
	mux.Get("/login", middlewares.LoggedIn(pages.Login))

	mux.Post("/register", middlewares.LoggedIn(procedures.Register))
	mux.Post("/login", middlewares.LoggedIn(procedures.Login))
	mux.Post("/logout", middlewares.LoggedIn(procedures.Logout))

	serverId := uuid.New("srv")
	s := server.NewServer(mux)

	s.Serve(url, serverId)
}
