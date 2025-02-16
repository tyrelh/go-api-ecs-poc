-- name: GetReward :one
SELECT * FROM rewards
WHERE id = ? LIMIT 1;

-- name: ListRewards :many
SELECT * FROM rewards
ORDER BY brand;

-- name: CreateReward :execresult
INSERT INTO rewards (
  brand, currency, denomination
) VALUES (
  ?, ?, ?
);

-- name: DeleteReward :exec
DELETE FROM rewards
WHERE id = ?;