package main

import (
	"context"
	"os"

	"scaffhold/config"
	"scaffhold/db/models"
	"scaffhold/handlers/pages"
	"scaffhold/handlers/procedures"
	"scaffhold/middlewares"
	pageTemplates "scaffhold/templates/pages"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/env"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func main() {
	// set up log.Logger
	err := logger.Setup("logs", "server-logs", config.LogLevel)
	if err != nil {
		os.Exit(1)
	}

	// set up env vars
	err = env.Load(&config.Vars)
	if err != nil {
		os.Exit(1)
	}

	// set up db
	db, err := database.Connect(config.Vars.DbPath)
	if err != nil {
		os.Exit(1)
	}

	// generate csrf tokens
	queries := models.New(db)
	sessionIDs, err := queries.ListSessionIDs(context.Background())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	err = csrf.Setup(sessionIDs)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// set up server
	url := server.NewLocalHostUrl(config.Port)

	router := server.NewRouter()

	router.Favicon("favicon.ico")

	router.Use(middlewares.Cache, middlewares.RequestID)

	router.TemplGet("/", pageTemplates.NotFound())
	router.Get("/{$}", pages.Index, middlewares.LoggedIn)
	router.Get("/register", pages.Register, middlewares.LoggedIn)
	router.Get("/login", pages.Login, middlewares.LoggedIn)

	router.Post("/register", procedures.Register, middlewares.LoggedIn)
	router.Post("/login", procedures.Login, middlewares.LoggedIn)
	router.Post("/logout", procedures.Logout, middlewares.LoggedIn)
	router.Post("/post", procedures.CreatePost, middlewares.LoggedIn, middlewares.CSRF)

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}