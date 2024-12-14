package helpers

import (
	"context"
	"log/slog"
	"scaffhold/db/models"
	"time"

	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/uuid"
)

func NewSession(queries *models.Queries, userId string) (string, error) {
	sessionId := uuid.New("ses")
	session, err := queries.CreateSession(context.Background(), models.CreateSessionParams{
		ID:         sessionId,
		UserID:     userId,
		ValidUntil: time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		return "", err
	}

	l.Logger.Debug("Session created", slog.String("user_id", userId))

	return session.ID, nil
}
