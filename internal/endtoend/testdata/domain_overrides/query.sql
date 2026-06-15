-- name: GetJob :one
SELECT * FROM jobs
WHERE id = $1 LIMIT 1;

-- name: ListJobsByStatus :many
SELECT * FROM jobs
WHERE status = $1
ORDER BY priority;
