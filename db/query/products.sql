-- name: CreateProduct :execresult
INSERT INTO
    products (name, price)
VALUES
    (?, ?);

-- name: GetProduct :one
SELECT
    *
FROM
    products
WHERE
    id = ?;

-- name: ListProducts :many
SELECT
    *
FROM
    products;

-- name: UpdateProduct :exec
UPDATE products
SET
    name = ?,
    price= ?
WHERE
    id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE
    id = ?;