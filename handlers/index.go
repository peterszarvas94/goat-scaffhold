package handlers

import (
	"net/http"
	"scaffhold/db/models"
	"scaffhold/templates/components"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if !ok || ctxUser == nil {
		server.TemplShow(pages.Index(&pages.IndexProps{
			LoginProps: &components.LoginProps{},
		}), w, r, http.StatusOK)
		return
	}

	server.TemplShow(pages.Index(&pages.IndexProps{
		UserinfoProps: &components.UserinfoProps{
			Name:  ctxUser.Name,
			Email: ctxUser.Email,
		},
	}), w, r, http.StatusOK)
}
