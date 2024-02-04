-- name: CreateMovie :one
INSERT INTO movies (title, genre) VALUES ($1, $2) RETURNING *;

-- name: GetMovies :many
SELECT * FROM movies;

-- name: GetMovie :one
SELECT * FROM movies WHERE id = $1 LIMIT 1;

-- name: UpdateMovie :one
UPDATE movies SET title=$2, genre=$3 WHERE id=$1 RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM movies WHERE id = $1;
