package procedures

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"scaffhold/db/models"
	"scaffhold/handlers/helpers"

	"github.com/peterszarvas94/goat/csrf"
	"github.com/peterszarvas94/goat/ctx"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	ctxUser, ok := ctx.GetFromCtx[models.User](r, "user")
	if ok && ctxUser != nil {
		// if logged in, redirect to index page
		helpers.HxRedirect(w, r, "/")
		return
	}

	if err := r.ParseForm(); err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	email := r.FormValue("email")
	if email == "" {
		helpers.BadRequest(w, r, "Email can not be empty")
		return
	}

	password := r.FormValue("password")
	if password == "" {
		helpers.BadRequest(w, r, "Password can not be empty")
		return
	}

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	queries := models.New(db)
	user, err := queries.Login(context.Background(), models.LoginParams{
		Email:    email,
		Password: password,
	})

	if err != nil {
		// wrong credentials
		helpers.Unauthorized(w, r, "Wrong credentials",
			slog.String("email", email),
			slog.String("password", password),
		)
		return
	}

	logger.AddToContext("user_id", user.ID)

	sessionId := uuid.New("ses")
	session, err := queries.CreateSession(context.Background(), models.CreateSessionParams{
		ID:         sessionId,
		UserID:     user.ID,
		ValidUntil: time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	logger.AddToContext("session_id", session.ID)

	csrfToken, err := csrf.AddNewCSRFToken(session.ID)
	if err != nil {
		helpers.ServerError(w, r, err)
		return
	}

	logger.AddToContext("csrf_token", csrfToken)

	helpers.SetCookie(&w, session.ID)

	logger.Debug("Logged in")

	helpers.HxRedirect(w, r, "/")
}
