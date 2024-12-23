package helpers

import (
	"net/http"
)

func ResetCookie(w *http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "sessionToken",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(*w, cookie)
}
