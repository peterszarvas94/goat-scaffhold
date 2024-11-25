package main

import (
	"context"
	"log/slog"
	"os"

	"bootstrap/config"
	"bootstrap/db/models"
	"bootstrap/handlers"
	"bootstrap/templates/pages"

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
	err = database.StartSqliteConnection(config.Vars.DbPath)
	if err != nil {
		os.Exit(1)
	}

	db, err := database.Get()

	if config.Vars.Env == "dev" {
		// seed
		queries := models.New(db)
		user, err := queries.CreateUser(context.Background(), models.CreateUserParams{
			ID:   uuid.New("usr"),
			Name: "Peter",
		})

		if err != nil {
			l.Logger.Error(err.Error())
		}

		l.Logger.Debug("user created", slog.String("id", user.ID), slog.String("name", user.Name))
	}

	// set up server
	url := server.NewLocalHostUrl(config.Port)

	mux := server.NewMux(url)
	list := []string{"one", "two", "three"}
	mux.TemplGet("/{$}", pages.Index.Full(list))
	mux.Get("/hello/{$}", handlers.MyHandlerFunc)

	serverId := uuid.New("srv")
	s := server.NewServer(mux)

	s.Serve(url, serverId)

}
