package handlers

import (
	"context"
	"net/http"

	"bootstrap/db/models"
	"bootstrap/templates/pages"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	db, err := database.Get()
	if err != nil {
		server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
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
		l.Logger.Error(err.Error())
		return
	}

	server.TemplShow(pages.Index(pages.IndexProps{
		User: &models.User{
			Name:  user.Name,
			Email: user.Email,
		},
		Partial: "userinfo",
	}), w, r, http.StatusOK)
}
