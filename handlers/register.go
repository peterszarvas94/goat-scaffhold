package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"scaffhold/db/models"
	"scaffhold/templates/components"

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

	existing, _ := queries.GetUserByEmail(context.Background(), email)
	var nameError string
	var emailError string

	if existing.Name == name {
		nameError = "Name already in use"
	}

	if existing.Email == email {
		emailError = "Email already in use"
	}

	if nameError != "" || emailError != "" {
		server.TemplShow(components.Register(components.RegisterProps{
			NameValue:     name,
			NameError:     nameError,
			EmailValue:    email,
			EmailError:    emailError,
			PasswordValue: password,
		}), w, r, http.StatusConflict)
		return
	}

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

	server.TemplShow(components.Login(components.LoginProps{
		EmailValue: user.Email,
	}), w, r, http.StatusOK)
}

func RegisterWidget(w http.ResponseWriter, r *http.Request) {
	l.Logger.Debug("Register widget requested")

	server.TemplShow(components.Register(components.RegisterProps{}), w, r, http.StatusOK)
}
