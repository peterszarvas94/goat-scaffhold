package helpers

import (
	"fmt"
	"net/http"

	l "github.com/peterszarvas94/goat/logger"
)

func HxRedirect(w http.ResponseWriter, r *http.Request, path string) {
	l.Logger.Debug(fmt.Sprintf("Redirecting to %s", path))
	w.Header().Set("HX-Redirect", path)
	w.WriteHeader(http.StatusMovedPermanently)
	w.Write([]byte{})
}
