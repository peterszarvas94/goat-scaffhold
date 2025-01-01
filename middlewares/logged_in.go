package middlewares

import (
	"context"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"time"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
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

			// no cookie
			logger.Debug("Cookie not found")
			next(w, r)
			return
		}

		logger.Add("session_id", cookie.Value)

		queries := models.New(db)
		session, err := queries.GetSessionByID(context.Background(), cookie.Value)
		if err != nil {
			helpers.ResetCookie(&w)

			csrf.DeleteCSRFToken(cookie.Value)

			// cookie, but no session -> next
			logger.Debug("Cookie is present, but no session is found")
			next(w, r)
			return
		}

		if session.ValidUntil.Before(time.Now()) {
			err = queries.DeleteSession(context.Background(), session.ID)
			if err != nil {
				helpers.ServerError(w, r, err)
				return
			}

			csrf.DeleteCSRFToken(session.ID)

			logger.Debug("Session is expired, deleted it")
			next(w, r)
			return
		}

		user, err := queries.GetUserByID(context.Background(), session.UserID)
		if err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		logger.Add("user_id", user.ID)

		items := ctx.KV{
			"user":    &user,
			"session": &session,
		}

		// cookie, session and csrf token, and valid -> next with ctx
		r = ctx.AddToContext(r, items)
		next(w, r)
	}
}
