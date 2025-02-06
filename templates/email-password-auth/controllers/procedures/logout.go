package procedures

import (
	"context"
	"net/http"
	"scaffhold/controllers/helpers"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	reqID, err := helpers.CheckReqID(r)
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	_, _, err = helpers.CheckLoggedIn(r)
	if err != nil {
		// if not logged in, redirect to index page
		logger.Debug("Not even logged in", "req_id", reqID)
		helpers.HxRedirect(w, r, "/", "req_id", reqID)
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	logger.Debug("Cookie found", "req_id", reqID, "session_id", cookie.Value)

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	err = queries.DeleteSession(context.Background(), cookie.Value)
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	helpers.ResetCookie(&w)

	csrf.DeleteCSRFToken(cookie.Value)

	logger.Debug("Logged out", "req_id", reqID)
	helpers.HxRedirect(w, r, "/", "req_id", reqID)
}
