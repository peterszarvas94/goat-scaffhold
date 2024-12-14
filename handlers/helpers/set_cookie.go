package helpers

import (
	"log/slog"
	"net/http"
	"scaffhold/config"

	l "github.com/peterszarvas94/goat/logger"
)

func SetCookie(w *http.ResponseWriter, sessionID string) {
	secure := true
	if config.Vars.Env == "dev" {
		secure = false
	}

	httponly := true
	if config.Vars.Env == "dev" {
		httponly = false
	}

	cookie := &http.Cookie{
		Name:     "sessionToken",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: httponly,
		SameSite: http.SameSiteLaxMode,
		Secure:   secure,
		MaxAge:   3600,
	}

	http.SetCookie(*w, cookie)

	l.Logger.Debug("Cookie is set", slog.String("session_id", sessionID))

}
