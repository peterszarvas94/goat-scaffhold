package procedures

import (
	"context"
	"errors"
	"net/http"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/request"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.Get[string](r, "req_id")
	if reqID == nil || !ok {
		request.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	cookie, err := r.Cookie("sessionToken")
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	logger.Debug("Cookie found", "req_id", reqID, "session_id", cookie.Value)

	db, err := database.Get()
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	err = queries.DeleteSession(context.Background(), cookie.Value)
	if err != nil {
		request.ServerError(w, r, err, "req_id", reqID)
		return
	}

	request.ResetCookie(&w, "sessionToken")

	csrf.Delete(cookie.Value)

	logger.Debug("Logged out", "req_id", reqID)
	request.HxRedirect(w, r, "/", "req_id", reqID)
}
