package middlewares

import (
	"net/http"

	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func RequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.New("req")
		logger.AddToContext("req_id", reqID)
		next(w, r)
	}
}
