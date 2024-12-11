package handlers

import (
	"net/http"
	"scaffhold/templates/pages"

	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func ServerError(err error, w http.ResponseWriter, r *http.Request) {
	l.Logger.Error(err.Error())
	server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
}
