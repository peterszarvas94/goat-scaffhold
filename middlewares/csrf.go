package middlewares

import (
	"context"
	"net/http"
	"scaffhold/handlers/helpers"
)

func CSRF(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		csrfToken := r.FormValue("csrf_token")
		if csrfToken == "" {
			helpers.BadRequest(w, r, "CSRF token missing")
			return
		}

		ctx := context.WithValue(r.Context(), "csrf_token", csrfToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
