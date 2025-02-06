package pages

import (
	"context"
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/db/models"
	"scaffhold/views/components"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	reqID, err := helpers.CheckReqID(r)
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	props := &pages.IndexProps{}

	ctxUser, ctxSession, err := helpers.CheckLoggedIn(r)
	if err != nil {
		logger.Debug("Rendering index", "req_id", reqID)
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	csrfToken, err := csrf.GetCSRFToken(ctxSession.ID)
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	posts, err := queries.GetPostsByUserId(context.Background(), ctxUser.ID)
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
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

	logger.Debug("Rendering index", "req_id", reqID)
	server.Render(w, r, pages.Index(props), http.StatusOK)
}
