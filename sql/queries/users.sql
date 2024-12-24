-- name: CreateUser :one
INSERT INTO users (
    username, hashed_password, 
    last_name, first_name, 
    permissions
)
VALUES(
    $1, $2, 
    $3, $4, 
    $5
)
RETURNING *;

-- name: GetHashedPassword :exec
SELECT hashed_password
FROM users
WHERE username = $1;

-- name: UpdateUserPermissions :exec
UPDATE users
SET permissions = $2, updated_at = NOW()
WHERE username = $1;