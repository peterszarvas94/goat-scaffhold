package pages

import (
	"net/http"
	"scaffhold/db/models"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/server"
)

func Register(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ok && ctxUser != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	server.TemplShow(pages.Register(), w, r, http.StatusOK)
}
