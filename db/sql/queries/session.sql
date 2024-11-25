-- name: GetSessionByID :one
SELECT *
FROM session
WHERE id = ?;

-- name: ListSessions :many
SELECT *
FROM session
ORDER BY name;

-- name: CreateSession :one
INSERT INTO session (id, user_id)
VALUES (?, ?)
RETURNING *;

-- name: UpdateSession :one
UPDATE session
SET user_id = ?
WHERE id = ?
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM session
WHERE id = ?;
