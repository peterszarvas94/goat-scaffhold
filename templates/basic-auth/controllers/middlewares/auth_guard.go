package middlewares

import (
	"net/http"
	"scaffhold/controllers/helpers"
)

func AuthGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := helpers.CheckAuthStatus(r)
		if err != nil {
			helpers.ServerError(w, r, err)
			return
		}

		next(w, r)
	}
}
