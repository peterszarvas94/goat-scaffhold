package pages

import (
	"net/http"
	"scaffhold/db/models"
	"scaffhold/templates/pages"

	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Register(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ok && ctxUser != nil {
		l.Logger.Debug("Redirecting to \"/\"")
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	server.Render(w, r, pages.Register(), http.StatusOK)
}
