package procedures

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"scaffhold/db/models"
	"scaffhold/handlers/helpers"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ok && ctxUser != nil {
		// if logged in, redirect to index page
		w.Header().Add("HX-Redirect", "/")
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	db, err := database.Get()
	if err != nil {
		helpers.HandleServerError(w, r, err)
		return
	}

	queries := models.New(db)

	existing, err := queries.GetUserByEmail(context.Background(), email)
	// user conflict
	if err == nil {
		if existing.Name == name {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Name already in use")
			return
		}

		if existing.Email == email {
			w.WriteHeader(http.StatusConflict)
			fmt.Fprintf(w, "Email already in use")
			return
		}

		helpers.HandleServerError(w, r, errors.New("Unexpected error"))
		return
	}

	user, err := queries.CreateUser(context.Background(), models.CreateUserParams{
		ID:       uuid.New("usr"),
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		helpers.HandleServerError(w, r, err)
		return
	}

	l.Logger.Debug("Registered", slog.String("user_id", user.ID))

	w.Header().Set("HX-Redirect", "/login")
}
