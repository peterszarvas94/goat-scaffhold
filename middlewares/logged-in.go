package middlewares

import (
	"context"
	"log/slog"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"time"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
)

func LoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := database.Get()
		if err != nil {
			helpers.HandleServerError(w, r, err)
			return
		}

		cookie, err := r.Cookie("sessionToken")
		if err != nil {
			l.Logger.Debug("Cookie not found")

			// 1. no cookie -> next with err
			next(w, r)
			return
		}

		l.Logger.Debug("Cookie is read", slog.String("cookie_name", cookie.Name))

		queries := models.New(db)
		session, err := queries.GetSessionByID(context.Background(), cookie.Value)
		if err != nil {
			l.Logger.Debug("No session found, reseting cookie", slog.String("session_id", cookie.Value))

			helpers.ResetCookie(&w)

			l.Logger.Debug("Cookie is reseted", slog.String("cookie_name", cookie.Name))

			// 2. cookie but no session -> next with err
			next(w, r)
			return
		}

		l.Logger.Debug("Session existst", slog.String("session_id", session.ID))

		if session.ValidUntil.Before(time.Now()) {
			l.Logger.Debug("Session is expired", slog.String("session_id", session.ID))

			err = queries.DeleteSession(context.Background(), session.ID)
			if err != nil {
				helpers.HandleServerError(w, r, err)
				return
			}

			// 3. cookie and session exists but expired -> next with err
			next(w, r)
			return
		}

		user, err := queries.GetUserByID(context.Background(), session.UserID)
		if err != nil {
			helpers.HandleServerError(w, r, err)
			return
		}

		l.Logger.Debug("User exists", slog.String("user_id", user.ID))

		// 4. cookie and session exists, and valid -> next with ctx
		ctx := context.WithValue(r.Context(), "user", &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
