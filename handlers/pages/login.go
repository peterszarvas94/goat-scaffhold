package pages

import (
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/server"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ok && ctxUser != nil {
		helpers.HttpRedirect(w, r, "/")
		return
	}

	server.Render(w, r, pages.Login(), http.StatusOK)
}
