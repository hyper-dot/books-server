-- name: GetAllUsers :one
SELECT * FROM users
WHERE id = $1;

-- name: AddUser :exec
INSERT INTO users (
    username, email, password, salt, refresh_token
) VALUES (
    $1, $2, $3, $4, $5
);
