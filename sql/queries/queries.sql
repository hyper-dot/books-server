-- name: GetAllUsers :one
SELECT * FROM users
WHERE id = $1;
