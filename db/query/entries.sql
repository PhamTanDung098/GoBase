-- name: CreateAEntries :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: updateEntries :one
UPDATE entries 
SET amount = $1
WHERE id = $2
RETURNING *;

-- name: deleteEntries :exec
DELETE FROM entries WHERE id = $1;

