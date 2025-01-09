-- name: CreateUser :exec
INSERT INTO users (
    username, hashed_password, 
    last_name, first_name, 
    permissions
)
VALUES(
    $1, $2, 
    $3, $4, 
    $5
);

-- name: GetAccount :one
SELECT *
FROM users
WHERE username = $1;

-- name: UpdateUserPermissions :exec
UPDATE users
SET permissions = $2, updated_at = NOW()
WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

-- name: UpdateUserInfo :exec
UPDATE users
SET hashed_password = $2, last_name = $3, first_name = $4
WHERE username = $1;

-- name: DeleteAllUsers :exec
DELETE FROM users;