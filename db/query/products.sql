-- name: CreateProduct :execresult
INSERT INTO products (name, price) VALUES (?, ?);

-- name: GetProduct :one
SELECT * FROM products WHERE id = ?;

-- name: ListProducts :many
SELECT * FROM products;