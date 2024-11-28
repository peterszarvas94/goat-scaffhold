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

func Index(w http.ResponseWriter, r *http.Request) {
	db, err := database.Get()
	if err != nil {
		server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		l.Logger.Debug("Cookie not found")

		server.TemplShow(pages.Index(&pages.IndexProps{
			LoginProps: &components.LoginProps{},
		}), w, r, http.StatusOK)

		return
	}

	l.Logger.Debug("Cookie is read", slog.String("cookie_name", cookie.Name))

	queries := models.New(db)
	session, err := queries.GetSessionByID(context.Background(), cookie.Value)
	if err != nil {
		l.Logger.Debug("No session found, reseting cookie", slog.String("session_id", cookie.Value))

		cookie = &http.Cookie{
			Name:   "sessionToken",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}

		http.SetCookie(w, cookie)

		l.Logger.Debug("Cookie is reseted", slog.String("cookie_name", cookie.Name))

		server.TemplShow(pages.Index(&pages.IndexProps{
			LoginProps: &components.LoginProps{},
		}), w, r, http.StatusOK)
		return
	}

	l.Logger.Debug("Session existst", slog.String("session_id", cookie.Value))

	user, err := queries.GetUserByID(context.Background(), session.UserID)
	if err != nil {
		l.Logger.Error(err.Error())

		return
	}

	l.Logger.Debug("User exists", slog.String("user_id", user.ID))

	server.TemplShow(pages.Index(&pages.IndexProps{
		UserinfoProps: &components.UserinfoProps{
			Name:  user.Name,
			Email: user.Email,
		},
	}), w, r, http.StatusOK)
}
