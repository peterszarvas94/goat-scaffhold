package helpers

import (
	"errors"
	"net/http"

	"github.com/peterszarvas94/goat/ctx"
)

func CheckCsrf(r *http.Request) (string, error) {
	csrfToken, ok := ctx.Get[string](r, "csrf_token")
	if !ok || csrfToken == nil {
		return "", errors.New("CSRF token is missing")
	}

	return *csrfToken, nil
}
