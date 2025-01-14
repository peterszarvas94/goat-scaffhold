package main

import (
	"context"
	"os"

	"scaffhold/config"
	"scaffhold/controllers/middlewares"
	"scaffhold/controllers/pages"
	"scaffhold/controllers/procedures"
	"scaffhold/db/models"
	pageViews "scaffhold/views/pages"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/env"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func main() {
	// set up logger
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
		os.Exit(1)
	}

	// set up server
	url := server.NewLocalHostUrl(config.Port)

	router := server.NewRouter()

	router.Favicon("favicon.ico")

	router.Static("/scripts/", "./scripts")
	router.Static("/styles/", "./styles")

	router.Use(middlewares.Cache, middlewares.AddReqID)

	router.TemplGet("/", pageViews.NotFound())
	router.Get("/{$}", pages.Index, middlewares.IsLoggedIn)
	router.Get("/register", pages.Register, middlewares.IsLoggedIn)
	router.Get("/login", pages.Login, middlewares.IsLoggedIn)

	router.Post("/register", procedures.Register, middlewares.IsLoggedIn)
	router.Post("/login", procedures.Login, middlewares.IsLoggedIn)
	router.Post("/logout", procedures.Logout, middlewares.IsLoggedIn)
	router.Post("/post", procedures.CreatePost, middlewares.IsLoggedIn, middlewares.ValidateCsrf)

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}
