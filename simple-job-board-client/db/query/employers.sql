-- name: ListEmployers :many
SELECT * from employers ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: GetEmployerById :one
SELECT * from employers WHERE id = $1;

-- name: CreateEmployer :one
INSERT INTO employers(name, description, social_media_links) VALUES($1, $2, $3) RETURNING *;