package helpers

import (
	"fmt"
	"net/http"

	l "github.com/peterszarvas94/goat/logger"
)

func HttpRedirect(w http.ResponseWriter, r *http.Request, path string) {
	l.Logger.Debug(fmt.Sprintf("Redirecting to %s", path))
	http.Redirect(w, r, path, http.StatusMovedPermanently)
}
