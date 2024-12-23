package middlewares

import (
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/logger"
)

func CSRF(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxSession, ok := ctx.GetFromCtx[models.Session](r, "session")
		if !ok || ctxSession == nil {
			// not logged in
			helpers.Unauthorized(w, r, "Not logged in")
			return
		}

		if err := r.ParseForm(); err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		csrfToken := r.FormValue("csrf_token")
		if csrfToken == "" {
			helpers.BadRequest(w, r, "CSRF token missing")
			return
		}

		err := csrf.ValidateCSRFToken(ctxSession.ID, csrfToken)
		if err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		logger.AddToContext("csrf_token", csrfToken)

		items := ctx.KV{"csrf_token": &csrfToken}
		r = ctx.AddToContext(r, items)
		next.ServeHTTP(w, r)
	}
}
