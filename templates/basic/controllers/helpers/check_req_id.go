package helpers

import (
	"errors"
	"net/http"

	"github.com/peterszarvas94/goat/ctx"
)

func CheckReqID(r *http.Request) (string, error) {
	reqID, ok := ctx.GetFromCtx[string](r, "req_id")
	if !ok && reqID == nil {
		return "", errors.New("Request ID is missing")
	}

	return *reqID, nil
}
