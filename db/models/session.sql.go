// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: session.sql

package models

import (
	"context"
	"time"
)

const createSession = `-- name: CreateSession :one
INSERT INTO session (id, user_id, valid_until)
VALUES (?, ?, ?)
RETURNING id, user_id, valid_until, created_at, updated_at
`

type CreateSessionParams struct {
	ID         string
	UserID     string
	ValidUntil time.Time
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession, arg.ID, arg.UserID, arg.ValidUntil)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ValidUntil,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM session
WHERE id = ?
`

func (q *Queries) DeleteSession(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteSession, id)
	return err
}

const getSessionByID = `-- name: GetSessionByID :one
SELECT id, user_id, valid_until, created_at, updated_at
FROM session
WHERE id = ?
`

func (q *Queries) GetSessionByID(ctx context.Context, id string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByID, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ValidUntil,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listSessionIDs = `-- name: ListSessionIDs :many
SELECT "id"
FROM session
`

func (q *Queries) ListSessionIDs(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listSessionIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSessions = `-- name: ListSessions :many
SELECT id, user_id, valid_until, created_at, updated_at
FROM session
`

func (q *Queries) ListSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, listSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.ValidUntil,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSession = `-- name: UpdateSession :one
UPDATE session
SET valid_until = ?
WHERE id = ?
RETURNING id, user_id, valid_until, created_at, updated_at
`

type UpdateSessionParams struct {
	ValidUntil time.Time
	ID         string
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, updateSession, arg.ValidUntil, arg.ID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ValidUntil,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
