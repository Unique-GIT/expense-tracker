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
    AND created_at >= NOW() 
        - sqlc.arg(days)::int * INTERVAL '1 day'
        - sqlc.arg(months)::int * INTERVAL '1 month'
        - sqlc.arg(years)::int * INTERVAL '1 year'
GROUP BY l.id;

-- name: GetTransactions :many
SELECT t.id,l.labelName,t.object_name,t.cost,t.created_at
FROM transactions as t
INNER JOIN labels as l
on l.id = t.label_id
WHERE t.user_id = $1;