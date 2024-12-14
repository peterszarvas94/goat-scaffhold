package helpers

import "net/http"

func ServeUnauthorized(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.WriteHeader(http.StatusUnauthorized)
	next.ServeHTTP(w, r)
}
