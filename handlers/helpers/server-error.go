package helpers

import (
	"net/http"
	"scaffhold/templates/pages"

	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func HandleServerError(w http.ResponseWriter, r *http.Request, err error) {
	l.Logger.Error(err.Error())
	server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
}
