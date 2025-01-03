package pages

import (
	"context"
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/db/models"
	"scaffhold/views/components"
	"scaffhold/views/pages"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
)

func Index(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.GetFromCtx[string](r, "req_id")
	if !ok && reqID == nil {
		helpers.ServerError(w, r, errors.New("Request id is missing"))
		return
	}

	props := &pages.IndexProps{}

	ctxSession, ok := ctx.GetFromCtx[models.Session](r, "session")
	if !ok || ctxSession == nil {
		logger.Debug("No session, rendering index", "req_id", *reqID)
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	ctxUser, ok := ctx.GetFromCtx[models.User](r, "user")
	if !ok || ctxUser == nil {
		logger.Debug("No user, rendering index", "req_id", *reqID)
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	csrfToken, err := csrf.GetCSRFToken(ctxSession.ID)
	if err != nil {
		logger.Debug("No token, rendering index", "req_id", *reqID)
		server.Render(w, r, pages.Index(props), http.StatusOK)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", *reqID)
		return
	}

	queries := models.New(db)
	posts, err := queries.GetPostsByUserId(context.Background(), ctxUser.ID)
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", *reqID)
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

	logger.Debug("Rendering index", "req_id", *reqID)
	server.Render(w, r, pages.Index(props), http.StatusOK)
}
