package procedures

import (
	"context"
	"net/http"
	"scaffhold/db/models"
	"scaffhold/handlers/helpers"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := ctx.GetFromCtx[models.User](r, "user")
	if !ok || ctxUser == nil {
		// if not logged in, redirect to index page
		helpers.HxRedirect(w, r, "/")
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	queries := models.New(db)
	err = queries.DeleteSession(context.Background(), cookie.Value)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.AddToContext("session_id", cookie.Value)

	helpers.ResetCookie(&w)

	csrf.DeleteCSRFToken(cookie.Value)

	logger.Debug("Logged out")

	helpers.HxRedirect(w, r, "/")
}
