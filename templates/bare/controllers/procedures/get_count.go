package procedures

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
)

var count = 0

func GetCount(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.Get[string](r, "req_id")
	if !ok || reqID == nil {
		request.ServerError(w, r, errors.New("Request ID is missing"))
	}

	logger.Debug("Count", "req_id", *reqID)
	w.Header().Set("Content-Type", "text/html")
	w.Write(fmt.Appendf([]byte{}, "%d", count))
}
