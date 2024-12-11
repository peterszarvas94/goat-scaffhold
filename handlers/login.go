package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"scaffhold/config"
	"scaffhold/db/models"
	"scaffhold/templates/components"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
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
		server.TemplShow(components.Login(components.LoginProps{
			EmailValue:    email,
			PasswordValue: password,
			PasswordError: "Email or password is incorrect",
		}), w, r, http.StatusForbidden)

		return
	}

	l.Logger.Debug("Login successful", slog.String("user_id", user.ID))

	sessionId := uuid.New("ses")
	session, err := queries.CreateSession(context.Background(), models.CreateSessionParams{
		ID:     sessionId,
		UserID: user.ID,
	})

	if err != nil {
		ServerError(err, w, r)
		return
	}

	l.Logger.Debug("Session created", slog.String("user_id", user.ID))

	secure := true
	if config.Vars.Env == "dev" {
		secure = false
	}

	httponly := true
	if config.Vars.Env == "dev" {
		httponly = false
	}

	cookie := &http.Cookie{
		Name:     "sessionToken",
		Value:    session.ID,
		Path:     "/",
		HttpOnly: httponly,
		SameSite: http.SameSiteLaxMode,
		Secure:   secure,
		MaxAge:   3600,
	}

	http.SetCookie(w, cookie)

	l.Logger.Debug("Cookie is set", slog.String("session_id", session.ID))

	l.Logger.Debug("Userinfo widget requested", slog.String("user_id", user.ID))

	server.TemplShow(components.Userinfo(components.UserinfoProps{
		Name:  user.Name,
		Email: user.Email,
	}), w, r, http.StatusOK)
}

func LoginWidget(w http.ResponseWriter, r *http.Request) {
	l.Logger.Debug("Login widget requested")

	server.TemplShow(components.Login(components.LoginProps{}), w, r, http.StatusOK)
}
