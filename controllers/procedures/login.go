package procedures

import (
	"context"
	"net/http"
	"time"

	"scaffhold/controllers/helpers"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqID, err := helpers.CheckReqID(r)
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	_, _, err = helpers.CheckLoggedIn(r)
	if err == nil {
		// already logged in
		helpers.HxRedirect(w, r, "/", "req_id", reqID)
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
	user, err := queries.Login(context.Background(), models.LoginParams{
		Email:    email,
		Password: password,
	})

	if err != nil {
		// wrong credentials
		helpers.Unauthorized(w, r, "Wrong credentials", "req_id", reqID)
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
