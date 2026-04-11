-- name: AddLabel :one
INSERT INTO labels(id,labelName,user_id)
VALUES(
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;