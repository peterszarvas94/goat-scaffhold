package pages

import (
	"context"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"scaffhold/templates/components"
	"scaffhold/templates/pages"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	props := &pages.IndexProps{}

	ctxSession, ok := ctx.GetFromCtx[models.Session](r, "session")
	if !ok || ctxSession == nil {
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	ctxUser, ok := ctx.GetFromCtx[models.User](r, "user")
	if !ok || ctxUser == nil {
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	csrfToken, err := csrf.GetCSRFToken(ctxSession.ID)
	if err != nil {
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
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

	props.UserinfoProps = &components.UserinfoProps{
		Name:  ctxUser.Name,
		Email: ctxUser.Email,
	}

	props.Posts = posts

	props.PostformProps = &components.PostformProps{
		CSRFToken: csrfToken,
		UserID:    ctxUser.ID,
	}

	server.Render(w, r, pages.Index(props), http.StatusOK)
}
