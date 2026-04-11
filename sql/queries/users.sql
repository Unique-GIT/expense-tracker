-- name: AddUser :one
INSERT INTO users(id,userName,userNumber)
VALUES(
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: ResetAllUsers :exec
DELETE FROM users;

-- name: GetUserByNumber :one
SELECT * FROM users
WHERE userNumber = $1;