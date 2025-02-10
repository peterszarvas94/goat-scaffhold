package procedures

import (
	"context"
	"errors"
	"net/http"
	"time"

	"scaffhold/controllers/helpers"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/hash"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqID, ok := ctx.Get[string](r, "req_id")
	if reqID == nil || !ok {
		helpers.ServerError(w, r, errors.New("Request ID is missing"))
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		helpers.BadRequest(w, r, "Email can not be empty", "req_id", reqID)
		return
	}

	password := r.FormValue("password")
	if password == "" {
		helpers.BadRequest(w, r, "Password can not be empty", "req_id", reqID)
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)
	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		helpers.Unauthorized(w, r, "User with this email not found", "req_id", reqID)
		return
	}

	valid := hash.VerifyPassword(password, user.Password)
	if !valid {
		helpers.Unauthorized(w, r, "Bad credentials", "req_id", reqID)
		return
	}

	logger.Debug("Credentials are valid", "req_id", reqID)

	sessionId := uuid.New("ses")
	session, err := queries.CreateSession(context.Background(), models.CreateSessionParams{
		ID:         sessionId,
		UserID:     user.ID,
		ValidUntil: time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	logger.Debug("New session", "req_id", reqID)

	_, err = csrf.AddNewCSRFToken(session.ID)
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	helpers.SetCookie(&w, session.ID)

	logger.Debug("Logged in", "req_id", reqID)
	helpers.HxRedirect(w, r, "/", "req_id", reqID)
}
