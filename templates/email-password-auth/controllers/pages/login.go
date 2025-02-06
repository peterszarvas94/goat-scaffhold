package pages

import (
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqID, err := helpers.CheckReqID(r)
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	_, _, err = helpers.CheckLoggedIn(r)
	if err == nil {
		logger.Debug("Already logged in", "req_id", reqID)
		helpers.HttpRedirect(w, r, "/", "req_id", reqID)
		return
	}

	logger.Debug("Rendering login", "req_id", reqID)
	server.Render(w, r, pages.Login(), http.StatusOK)
}
