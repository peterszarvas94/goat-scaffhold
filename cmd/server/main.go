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

	router := server.NewRouter()

	router.Favicon("favicon.ico")

	router.Use(middlewares.Cache)

	router.TemplGet("/", pageTemplates.NotFound())
	router.Get("/{$}", middlewares.LoggedIn(pages.Index))
	router.Get("/register", middlewares.LoggedIn(pages.Register))
	router.Get("/login", middlewares.LoggedIn(pages.Login))

	router.Post("/register", middlewares.LoggedIn(procedures.Register))
	router.Post("/login", middlewares.LoggedIn(procedures.Login))
	router.Post("/logout", middlewares.LoggedIn(procedures.Logout))
	router.Post("/post", middlewares.LoggedIn(middlewares.CSRF(procedures.CreatePost)))

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}
