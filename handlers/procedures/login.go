package procedures

import (
	"context"
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
		l.Logger.Debug("Redirecting to \"/\"")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		helpers.BadRequest(w, r, "Email can not be empty")
		return
	}

	password := r.FormValue("password")
	if password == "" {
		helpers.BadRequest(w, r, "Password can not be empty")
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	queries := models.New(db)
	user, err := queries.Login(context.Background(), models.LoginParams{
		Email:    email,
		Password: password,
	})

	if err != nil {
		// wrong credentials
		helpers.Unauthorized(w, r, "Wrong credentials",
			slog.String("email", email),
			slog.String("password", password),
		)
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
		helpers.ServerError(w, r, err)
		return
	}

	l.Logger.Debug("Session created", slog.String("user_id", user.ID))

	helpers.SetCookie(&w, session.ID)

	l.Logger.Debug("Redirecting to \"/\"")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
