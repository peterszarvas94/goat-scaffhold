package pages

import (
	"net/http"
	"scaffhold/db/models"
	"scaffhold/templates/components"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var props *pages.IndexProps

	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ok && ctxUser != nil {
		props = &pages.IndexProps{
			UserinfoProps: &components.UserinfoProps{
				Name:  ctxUser.Name,
				Email: ctxUser.Email,
			},
		}
	} else {
		props = &pages.IndexProps{}
	}

	server.Render(w, r, pages.Index(props), http.StatusOK)
}
