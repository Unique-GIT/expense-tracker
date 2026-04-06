-- name: AddRecord :one
INSERT INTO transactions(id, object, cost, user_id, label_id)
VALUES(
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;