package pages

import (
	"context"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"scaffhold/templates/components"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	props := &pages.IndexProps{}

	ctxUser, ok := r.Context().Value("user").(*models.User)
	if !ok || ctxUser == nil {
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	props.UserinfoProps = &components.UserinfoProps{
		Name:  ctxUser.Name,
		Email: ctxUser.Email,
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	queries := models.New(db)
	posts, err := queries.GetPostsByUserId(context.Background(), ctxUser.ID)
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	props.Posts = posts

	props.PostformProps = &components.PostformProps{
		CSRFToken: "very_good_token",
		UserID:    ctxUser.ID,
	}

	server.Render(w, r, pages.Index(props), http.StatusOK)
}
