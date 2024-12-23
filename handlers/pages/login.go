package pages

import (
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := ctx.GetFromCtx[models.User](r, "user")
	if ok && ctxUser != nil {
		helpers.HttpRedirect(w, r, "/")
		return
	}

	logger.Debug("Rendering login")
	server.Render(w, r, pages.Login(), http.StatusOK)
}
