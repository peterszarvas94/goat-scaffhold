package middlewares

import (
	"errors"
	"net/http"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/request"
)

func CSRFGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctxSession, ok := ctx.Get[models.Session](r, "session")
		if !ok || ctxSession == nil {
			request.Unauthorized(w, r, errors.New("Not logged in"))
			return
		}

		err := r.ParseForm()
		if err != nil {
			request.ServerError(w, r, err)
			return
		}

		csrfToken := r.FormValue("csrf_token")

		err = csrf.Validate(ctxSession.ID, csrfToken)
		if err != nil {
			request.ServerError(w, r, errors.New("CSRF token is invalid"))
			return
		}

		next(w, r)
	}
}
