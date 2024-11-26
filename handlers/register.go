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
	"github.com/peterszarvas94/goat/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		ServerError(err, w, r)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	db, err := database.Get()
	if err != nil {
		ServerError(err, w, r)
		return
	}

	queries := models.New(db)
	user, err := queries.CreateUser(context.Background(), models.CreateUserParams{
		ID:       uuid.New("usr"),
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		ServerError(err, w, r)
		return
	}

	l.Logger.Debug("Registered", slog.String("user_id", user.ID))

	LoginWidget(w, r)
}

func RegisterWidget(w http.ResponseWriter, r *http.Request) {
	l.Logger.Debug("Register widget requested")

	server.TemplShow(components.Register(), w, r, http.StatusOK)
}
