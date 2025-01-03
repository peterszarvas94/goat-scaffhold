package pages

import (
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/db/models"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Register(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.GetFromCtx[string](r, "req_id")
	if !ok && reqID == nil {
		helpers.ServerError(w, r, errors.New("Request id is missing"))
		return
	}

	ctxUser, ok := ctx.GetFromCtx[models.User](r, "user")
	if ok && ctxUser != nil {
		logger.Debug("Already logged in", "req_id", *reqID)
		helpers.HttpRedirect(w, r, "/", "req_id", *reqID)
		return
	}

	logger.Debug("Rendering Register", "req_id", *reqID)
	server.Render(w, r, pages.Register(), http.StatusOK)
}
