-- name: CreateProduct :execresult
INSERT INTO
    products (name, price, is_active, sale_price, sku)
VALUES
    (?, ?, ?, ?, ?);

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
    price = ?,
    is_active = ?,
    sale_price = ?,
    sku = ?
WHERE
    id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE
    id = ?;