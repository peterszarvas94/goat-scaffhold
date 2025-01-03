package helpers

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func HttpRedirect(w http.ResponseWriter, r *http.Request, path string, args ...any) {
	logger.Debug(fmt.Sprintf("Redirecting to %s", path), args...)
	http.Redirect(w, r, path, http.StatusMovedPermanently)
}
