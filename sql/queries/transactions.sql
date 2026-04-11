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

-- name: GetAnalysis :many
SELECT l.labelName,
    SUM(t.cost)::REAL AS total_cost
FROM transactions as t
INNER JOIN labels as l
ON l.id = t.label_id
WHERE t.user_id = $1
GROUP BY l.id;