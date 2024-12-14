package helpers

import (
	"log/slog"
	"net/http"

	l "github.com/peterszarvas94/goat/logger"
)

func ResetCookie(w *http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessionToken",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(*w, cookie)
	l.Logger.Debug("Cookie is reset", slog.String("cookie_name", cookie.Name))
}
