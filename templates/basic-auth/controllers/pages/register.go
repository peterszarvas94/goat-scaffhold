package pages

import (
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Register(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.Get[string](r, "req_id")
	if reqID == nil || !ok {
		helpers.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	err := helpers.CheckGuestStatus(r)
	if err != nil {
		logger.Debug(err.Error(), "req_id", *reqID)
		helpers.HttpRedirect(w, r, "/", "req_id", *reqID)
		return
	}

	logger.Debug("Rendering register page", "req_id", *reqID)
	server.Render(w, r, pages.Register(), http.StatusOK)
}
