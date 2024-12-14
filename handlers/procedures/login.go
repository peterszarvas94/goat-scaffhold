package procedures

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"scaffhold/db/models"
	"scaffhold/handlers/helpers"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ok && ctxUser != nil {
		// if logged in, redirect to index page
		w.Header().Add("HX-Redirect", "/")
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.HandleServerError(w, r, err)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	db, err := database.Get()
	if err != nil {
		helpers.HandleServerError(w, r, err)
		return
	}

	queries := models.New(db)
	user, err := queries.Login(context.Background(), models.LoginParams{
		Email:    email,
		Password: password,
	})

	if err != nil {
		// wrong credentials
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Email or password is incorrect")
		return
	}

	// correct credentials
	l.Logger.Debug("Login successful", slog.String("user_id", user.ID))

	sessionId := uuid.New("ses")
	session, err := queries.CreateSession(context.Background(), models.CreateSessionParams{
		ID:         sessionId,
		UserID:     user.ID,
		ValidUntil: time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		helpers.HandleServerError(w, r, err)
		return
	}

	l.Logger.Debug("Session created", slog.String("user_id", user.ID))

	helpers.SetCookie(&w, session.ID)

	w.Header().Set("HX-Redirect", "/")
}
