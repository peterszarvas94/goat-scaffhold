package handlers

import (
	"net/http"
	"scaffhold/templates/pages"

	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func BadRequest(err error, w http.ResponseWriter, r *http.Request) {
	l.Logger.Error(err.Error())
	server.TemplShow(pages.BadRequest(), w, r, http.StatusBadRequest)
}
