-- name: AddLabel :one
INSERT INTO labels(id,labelName,user_id)
VALUES(
    gen_random_uuid(),
    $1,
    $2
)
RETURNING *;

-- name: GetLabels :many
SELECT labelName FROM labels
WHERE user_id = $1;