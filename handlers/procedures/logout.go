package procedures

import (
	"context"
	"log/slog"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if !ok || ctxUser == nil {
		// if not logged in, redirect to index page
		l.Logger.Debug("Redirecting to \"/\"")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		l.Logger.Error(err.Error())
		return
	}

	l.Logger.Debug("Cookie is read", slog.String("cookie_name", cookie.Name))

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	queries := models.New(db)
	err = queries.DeleteSession(context.Background(), cookie.Value)
	if err != nil {
		l.Logger.Error(err.Error())
		return
	}

	l.Logger.Debug("Session is deleted", slog.String("session_id", cookie.Value))

	helpers.ResetCookie(&w)

	l.Logger.Debug("Cookie is invalidated", slog.String("cookie_name", cookie.Name))

	l.Logger.Debug("Logged out")

	l.Logger.Debug("Redirecting to \"/\"")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
