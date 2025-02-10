package middlewares

import (
	"net/http"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/uuid"
)

func AddRequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.New("req")

		items := ctx.KV{
			"req_id": &reqID,
		}

		r = ctx.Add(r, items)
		next(w, r)
	}
}
