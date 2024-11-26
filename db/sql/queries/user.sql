-- name: GetUserByID :one
SELECT *
FROM user
WHERE id = ?;

-- name: ListUser :many
SELECT *
FROM user
ORDER BY name;

-- name: CreateUser :one
INSERT INTO user (id, name, email, password)
VALUES (?, ?, ?, ?)
RETURNING id, name, email;

-- name: Login :one
SELECT id, email, name
FROM user
WHERE email = ? AND password = ?;

-- name: UpdateUser :one
UPDATE user
SET name = ?, email = ?, password = ?
WHERE id = ?
RETURNING id, name, email;

-- name: DeleteUser :exec
DELETE FROM user
WHERE id = ?;
