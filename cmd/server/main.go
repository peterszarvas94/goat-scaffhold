package main

import (
	"os"

	"scaffhold/config"
	"scaffhold/handlers"
	"scaffhold/middlewares"
	"scaffhold/templates/pages"

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
	mux.TemplGet("/", pages.NotFound())

	mux.Get("/{$}", middlewares.LoggedIn(handlers.Index))

	mux.Get("/register", handlers.RegisterWidget)
	mux.Post("/register", middlewares.LoggedIn(handlers.Register))

	mux.Get("/login", handlers.LoginWidget)
	mux.Post("/login", middlewares.LoggedIn(handlers.Login))

	mux.Post("/logout", middlewares.LoggedIn(handlers.Logout))

	serverId := uuid.New("srv")
	s := server.NewServer(mux)

	s.Serve(url, serverId)
}
