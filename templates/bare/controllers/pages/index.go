package pages

import (
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	reqID, err := helpers.CheckReqID(r)
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	logger.Debug("Rendering index", "req_id", reqID)
	server.Render(w, r, pages.Index(), http.StatusOK)
}
