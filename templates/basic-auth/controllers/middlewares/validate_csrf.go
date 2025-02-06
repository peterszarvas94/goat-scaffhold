package middlewares

import (
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
)

func ValidateCsrf(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ctxSession, err := helpers.CheckLoggedIn(r)
		if err != nil {
			helpers.Unauthorized(w, r, "Not logged in")
			return
		}

		if err = r.ParseForm(); err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		csrfToken := r.FormValue("csrf_token")

		err = csrf.ValidateCSRFToken(ctxSession.ID, csrfToken)
		if err != nil {
			helpers.ServerError(w, r, errors.New("CSRF token is invalid"))
			return
		}

		items := ctx.KV{
			"csrf_token": &csrfToken,
		}

		r = ctx.AddToContext(r, items)
		next.ServeHTTP(w, r)
	}
}
