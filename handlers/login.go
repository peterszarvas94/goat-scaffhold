package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"bootstrap/db/models"
	"bootstrap/templates/components"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ServerError(err, w, r)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	db, err := database.Get()
	if err != nil {
		ServerError(err, w, r)
		return
	}

	queries := models.New(db)
	user, err := queries.Login(context.Background(), models.LoginParams{
		Email:    email,
		Password: password,
	})

	if err != nil {
		ServerError(err, w, r)
		return
	}

	l.Logger.Debug("Logged in", slog.String("user_id", user.ID))

	l.Logger.Debug("Userinfo widget requested", slog.String("user_id", user.ID))

	server.TemplShow(components.Userinfo(&models.User{
		Name:  user.Name,
		Email: user.Email,
	}), w, r, http.StatusOK)
}

func LoginWidget(w http.ResponseWriter, r *http.Request) {
	l.Logger.Debug("Login widget requested")

	server.TemplShow(components.Login(), w, r, http.StatusOK)
}
