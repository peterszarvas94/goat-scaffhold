package helpers

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func Conflict(w http.ResponseWriter, r *http.Request, msg string, args ...any) {
	logger.Error(msg, args...)
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintln(w, msg)
}
