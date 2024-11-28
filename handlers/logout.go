package handlers

import (
	"bootstrap/db/models"
	"bootstrap/templates/components"
	"bootstrap/templates/pages"
	"context"
	"log/slog"
	"net/http"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	db, err := database.Get()
	if err != nil {
		server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		l.Logger.Error(err.Error())
		return
	}

	l.Logger.Debug("Cookie is read", slog.String("cookie_name", cookie.Name))

	queries := models.New(db)
	err = queries.DeleteSession(context.Background(), cookie.Value)
	if err != nil {
		l.Logger.Error(err.Error())
		return
	}

	l.Logger.Debug("Session is deleted", slog.String("session_id", cookie.Value))

	cookie = &http.Cookie{
		Name:   "sessionToken",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	l.Logger.Debug("Cookie is invalidated", slog.String("cookie_name", cookie.Name))

	l.Logger.Debug("Logged out")

	server.TemplShow(components.Login(components.LoginProps{}), w, r, http.StatusOK)
}
