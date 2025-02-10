package procedures

import (
	"context"
	"errors"
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/db/models"
	"scaffhold/views/components"

	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.Get[string](r, "req_id")
	if reqID == nil || !ok {
		helpers.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	ctxUser, _, err := helpers.CheckAuthStatus(r)
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", *reqID)
		return
	}

	_, err = helpers.CheckCsrf(r)
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", *reqID)
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.ServerError(w, r, err, "req_id", *reqID)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		helpers.HxRe(w, "#post-error", "innerHtml")
		helpers.BadRequest(w, r, "Title can not be empty", "req_id", *reqID)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		helpers.HxRe(w, "#post-error", "innerHtml")
		helpers.BadRequest(w, r, "Content can not be empty", "req_id", *reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", *reqID)
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
		helpers.ServerError(w, r, err, "req_id", *reqID)
		return
	}

	logger.Debug("Post created, rendering new post", "req_id", *reqID)

	server.Render(w, r, components.Post(&models.Post{
		Title: post.Title,
	}), http.StatusOK)
}
