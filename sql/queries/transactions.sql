-- name: AddTransaction :one
INSERT INTO transactions(id,user_id,label_id,object_name,cost,created_at)
VALUES(
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    NOW()
)
RETURNING *;