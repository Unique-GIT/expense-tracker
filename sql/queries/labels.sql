-- name: GetLabels :many
SELECT id,name FROM labels;

-- name: GetLabelByName :one
SELECT id,name FROM labels
WHERE name=$1;