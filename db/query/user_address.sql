-- name: CreateUserAddress :execresult
INSERT INTO
    user_addresses (user_id, street, city, country)
VALUES
    (?, ?, ?, ?);

-- name: GetAddressesByUserID :many
SELECT
    *
FROM
    user_addresses
WHERE
    user_id = ?;

-- name: UpdateUserAddress :exec
UPDATE user_addresses
SET
    user_id = ?,
    street = ?,
    city = ?,
    country = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND user_id = ?;

-- name: GetAllUserAddresses :many
SELECT
    *
FROM
    user_addresses;

-- name: DeleteAddressesByUserID :exec
DELETE FROM user_addresses
WHERE
    user_id = ?;