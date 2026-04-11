-- name: AddUser :one
INSERT INTO users(id,userName,userNumber)
VALUES(
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;