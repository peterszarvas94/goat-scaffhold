package procedures

import (
	"context"
	"errors"
	"net/http"

	"scaffhold/controllers/helpers"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
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

	name := r.FormValue("name")
	if name == "" {
		helpers.BadRequest(w, r, "Name can not be empty", "req_id", reqID)
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

	// TODO: has password

	db, err := database.Get()
	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	queries := models.New(db)

	existing, err := queries.GetUserByEmail(context.Background(), email)
	// user conflict
	if err == nil {
		if existing.Name == name {
			helpers.Conflict(w, r, "Name already in use", "req_id", reqID)
			return
		}

		if existing.Email == email {
			helpers.Conflict(w, r, "Email already in use", "req_id", reqID)
			return
		}

		helpers.ServerError(w, r, errors.New("Conflict"), "req_id", reqID)
		return
	}

	_, err = queries.CreateUser(context.Background(), models.CreateUserParams{
		ID:       uuid.New("usr"),
		Name:     name,
		Email:    email,
		Password: password,
	})

	if err != nil {
		helpers.ServerError(w, r, err, "req_id", reqID)
		return
	}

	logger.Debug("Registered", "req_id", reqID)
	helpers.HxRedirect(w, r, "/login", "req_id", reqID)
}
