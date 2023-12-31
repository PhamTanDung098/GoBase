// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: entries.sql

package db

import (
	"context"
)

const createAEntries = `-- name: CreateAEntries :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING id, account_id, amount, created_at
`

type CreateAEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateAEntries(ctx context.Context, arg CreateAEntriesParams) (Entries, error) {
	row := q.db.QueryRowContext(ctx, createAEntries, arg.AccountID, arg.Amount)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntries = `-- name: DeleteEntries :exec
DELETE FROM entries WHERE id = $1
`

func (q *Queries) DeleteEntries(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntries, id)
	return err
}

const getEntries = `-- name: GetEntries :one
SELECT id, account_id, amount, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntries(ctx context.Context, id int64) (Entries, error) {
	row := q.db.QueryRowContext(ctx, getEntries, id)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, account_id, amount, created_at FROM entries
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListEntriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entries{}
	for rows.Next() {
		var i Entries
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntries = `-- name: UpdateEntries :one
UPDATE entries 
SET amount = $1
WHERE id = $2
RETURNING id, account_id, amount, created_at
`

type UpdateEntriesParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateEntries(ctx context.Context, arg UpdateEntriesParams) (Entries, error) {
	row := q.db.QueryRowContext(ctx, updateEntries, arg.Amount, arg.ID)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
