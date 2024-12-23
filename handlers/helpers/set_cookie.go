package helpers

import (
	"net/http"
	"scaffhold/config"
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
}
