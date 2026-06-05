-- name: CreateUser :execresult
INSERT INTO
    users (name, email)
VALUES
    (?, ?);

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = ?;

-- name: ListUsers :many
SELECT
    *
FROM
    users;

-- name: UpdateUser :exec
UPDATE users
SET
    name = ?,
    email = ?
WHERE
    id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = ?;

-- name: UpdateAvatar :exec
UPDATE users
SET
    avatar_url = ?
WHERE
    id = ?;

-- name: