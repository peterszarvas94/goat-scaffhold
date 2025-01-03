package middlewares

import (
	"context"
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/db/models"
	"time"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
)

func LoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID, ok := ctx.GetFromCtx[string](r, "req_id")
		if !ok && reqID == nil {
			helpers.ServerError(w, r, errors.New("Request id is missing"))
			return
		}

		db, err := database.Get()
		if err != nil {
			helpers.ServerError(w, r, err, "req_id", *reqID)
			return
		}

		cookie, err := r.Cookie("sessionToken")
		if err != nil {

			// no cookie
			logger.Debug("Cookie not found", "req_id", *reqID)
			next(w, r)
			return
		}

		logger.Debug("Cookie found", "req_id", *reqID)

		queries := models.New(db)
		session, err := queries.GetSessionByID(context.Background(), cookie.Value)
		if err != nil {
			helpers.ResetCookie(&w)

			csrf.DeleteCSRFToken(cookie.Value)

			// cookie, but no session -> next
			logger.Debug("Cookie is found, but no session", "req_id", *reqID)
			return
		}

		if session.ValidUntil.Before(time.Now()) {
			logger.Debug("Session is expired", "req_id", *reqID, "session_id", session.ID)

			err = queries.DeleteSession(context.Background(), session.ID)
			if err != nil {
				helpers.ServerError(w, r, err, "req_id", *reqID)
				return
			}

			csrf.DeleteCSRFToken(session.ID)
			next(w, r)
			return
		}

		logger.Debug("Session valid", "req_id", *reqID, "session_id", session.ID)

		user, err := queries.GetUserByID(context.Background(), session.UserID)
		if err != nil {
			helpers.ServerError(w, r, err, "req_id", *reqID)
			return
		}

		logger.Debug("User exist", "req_id", *reqID, "user_id", user.ID)

		items := ctx.KV{
			"user":    &user,
			"session": &session,
		}

		// cookie, session and csrf token, and valid -> next with ctx
		r = ctx.AddToContext(r, items)
		next(w, r)
	}
}
