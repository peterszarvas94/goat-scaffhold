package middlewares

import (
	"context"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
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

		ctx := context.WithValue(r.Context(), "csrf_token", csrfToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
