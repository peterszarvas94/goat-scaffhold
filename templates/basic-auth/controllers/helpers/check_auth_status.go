package helpers

import (
	"errors"
	"net/http"
	"scaffhold/db/models"

	"github.com/peterszarvas94/goat/ctx"
)

func CheckAuthStatus(r *http.Request) (*models.User, *models.Session, error) {
	session, ok := ctx.Get[models.Session](r, "session")
	if !ok || session == nil {
		return nil, nil, errors.New("Session is missing")
	}

	user, ok := ctx.Get[models.User](r, "user")
	if !ok || user == nil {
		return nil, nil, errors.New("User is missing")
	}

	return user, session, nil
}
