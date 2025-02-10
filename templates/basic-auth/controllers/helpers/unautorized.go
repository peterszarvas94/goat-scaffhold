package helpers

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, err error, args ...any) {
	logger.Error(err.Error(), args...)
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w, "Unauthorized")
}
