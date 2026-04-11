-- name: AddTransaction :one
INSERT INTO transactions(id,user_id,label_id,object_name,cost)
VALUES(
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;