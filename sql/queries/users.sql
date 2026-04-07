-- name: GetUserByName :one
SELECT id,name,number FROM users
WHERE name = $1;