package helpers

import (
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/logger"
)

func HttpRedirect(w http.ResponseWriter, r *http.Request, path string) {
	logger.Debug(fmt.Sprintf("Redirecting to %s", path))
	http.Redirect(w, r, path, http.StatusMovedPermanently)
}
