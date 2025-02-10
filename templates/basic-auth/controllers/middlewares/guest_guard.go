package middlewares

import (
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"
)

func GuestGuard(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := helpers.CheckAuthStatus(r)
		if err == nil {
			helpers.ServerError(w, r, errors.New("User should not be logged in"))
			return
		}
		next(w, r)
	}
}
