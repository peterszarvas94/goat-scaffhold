package helpers

import (
	"errors"
	"net/http"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/ctx"
)

func CheckLoggedIn(r *http.Request) (*models.User, *models.Session, error) {
	session, ok := ctx.GetFromCtx[models.Session](r, "session")
	if !ok || session == nil {
		return nil, nil, errors.New("Session is missing")
	}

	user, ok := ctx.GetFromCtx[models.User](r, "user")
	if !ok || user == nil {
		return nil, nil, errors.New("User is missing")
	}

	return user, session, nil
}
