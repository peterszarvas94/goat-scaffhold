package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"time"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
)

func LoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := database.Get()
		if err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		cookie, err := r.Cookie("sessionToken")
		if err != nil {
			l.Logger.Debug("Cookie not found")

			// no cookie
			next(w, r)
			return
		}

		l.Logger.Debug("Cookie is read", slog.String("cookie_name", cookie.Name))

		queries := models.New(db)
		session, err := queries.GetSessionByID(context.Background(), cookie.Value)
		if err != nil {
			l.Logger.Debug("No session found", slog.String("session_id", cookie.Value))

			helpers.ResetCookie(&w)

			l.Logger.Debug("Cookie is reseted", slog.String("cookie_name", cookie.Name))

			csrf.DeleteCSRFToken(cookie.Value)

			// cookie, but no session -> next
			next(w, r)
			return
		}

		l.Logger.Debug("Session existst", slog.String("session_id", session.ID))

		if session.ValidUntil.Before(time.Now()) {
			l.Logger.Debug("Session is expired", slog.String("session_id", session.ID))

			err = queries.DeleteSession(context.Background(), session.ID)
			if err != nil {
				helpers.ServerError(w, r, err)
				return
			}

			csrf.DeleteCSRFToken(session.ID)

			// cookie and session, but expired -> next
			next(w, r)
			return
		}

		user, err := queries.GetUserByID(context.Background(), session.UserID)
		if err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		l.Logger.Debug("User exists", slog.String("user_id", user.ID))

		items := []ctx.KV{
			{Key: "user", Value: &user},
			{Key: "session", Value: &session},
		}

		// cookie, session and csrf token, and valid -> next with ctx
		r = ctx.AddToContext(r, items)
		next(w, r)
	}
}
