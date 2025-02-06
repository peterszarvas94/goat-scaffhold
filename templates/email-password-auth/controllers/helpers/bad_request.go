package helpers

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func BadRequest(w http.ResponseWriter, r *http.Request, msg string, args ...any) {
	logger.Warn(msg, args...)
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, msg)
}
