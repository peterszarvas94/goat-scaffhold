package procedures

import (
	"errors"
	"net/http"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
)

func PostCount(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.Get[string](r, "req_id")
	if !ok || reqID == nil {
		request.ServerError(w, r, errors.New("Request ID is missing"))
	}

	count++

	logger.Debug("Count increased", "req_id", *reqID)
	GetCount(w, r)
}
