package procedures

import (
	"context"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"scaffhold/templates/components"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := ctx.GetFromCtx[models.User](r, "user")
	if !ok || ctxUser == nil {
		helpers.Unauthorized(w, r, "Not logged in")
		return
	}

	token, ok := ctx.GetFromCtx[string](r, "csrf_token")
	if !ok || *token == "" {
		// token not found
		helpers.Unauthorized(w, r, "CSRF token invalid")
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		helpers.BadRequest(w, r, "Title can not be empty")
		return
	}

	content := r.FormValue("content")
	if content == "" {
		helpers.BadRequest(w, r, "Content can not be empty")
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	postId := uuid.New("pst")
	queries := models.New(db)
	post, err := queries.CreatePost(context.Background(), models.CreatePostParams{
		ID:      postId,
		Title:   title,
		Content: content,
		UserID:  ctxUser.ID,
	})

	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	logger.AddToContext("post_id", post.ID)
	logger.Debug("Post created")

	server.Render(w, r, components.Post(&models.Post{
		Title: post.Title,
	}), http.StatusOK)
}
