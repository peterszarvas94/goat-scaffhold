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
		l.Logger.Debug("Redirecting to \"/\"")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Name is empty")
		return
	}

	email := r.FormValue("email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Email is empty")
		return
	}

	password := r.FormValue("password")
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Password is empty")
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err)
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

		helpers.ServerError(w, r, errors.New("Unexpected error"))
		return
	}

	user, err := queries.CreateUser(context.Background(), models.CreateUserParams{
		ID:       uuid.New("usr"),
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	l.Logger.Debug("Registered", slog.String("user_id", user.ID))

	l.Logger.Debug("Redirecting to \"/\"")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
