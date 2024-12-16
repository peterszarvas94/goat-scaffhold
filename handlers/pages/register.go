package pages

import (
	"fmt"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/server"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in register")
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if ok && ctxUser != nil {
		helpers.HttpRedirect(w, r, "/")
		return
	}

	fmt.Println("rendering register")
	server.Render(w, r, pages.Register(), http.StatusOK)
}
