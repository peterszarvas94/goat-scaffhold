package procedures

import (
	"errors"
	"net/http"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.Get[string](r, "req_id")
	if !ok || reqID == nil {
		request.ServerError(w, r, errors.New("Request ID is missing"))
	}

	logger.Debug("Pong", "req_id", *reqID)
	w.Write([]byte("Pong "))
}
