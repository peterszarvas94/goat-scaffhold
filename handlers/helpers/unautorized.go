package helpers

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, msg string, args ...any) {
	logger.Error(msg, args...)
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w, msg)
}
