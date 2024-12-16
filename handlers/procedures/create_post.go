package procedures

import (
	"context"
	"log/slog"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"
	"scaffhold/templates/components"

	"github.com/peterszarvas94/goat/database"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := r.Context().Value("user").(*models.User)
	if !ok || ctxUser == nil {
		helpers.Unauthorized(w, r, "Not logged in")
		return
	}

	token, ok := r.Context().Value("csrf_token").(string)
	if !ok || token == "" {
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

	l.Logger.Debug("Post created", slog.String("post_id", post.ID))

	server.Render(w, r, components.Post(&models.Post{
		Title: post.Title,
	}), http.StatusOK)
}
