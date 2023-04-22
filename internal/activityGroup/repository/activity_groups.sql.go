// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: activity_groups.sql

package querier

import (
	"context"
	"database/sql"
)

const createActivityGroup = `-- name: CreateActivityGroup :execresult
INSERT INTO activity_groups (title, email)
VALUES (?, ?)
`

type CreateActivityGroupParams struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

func (q *Queries) CreateActivityGroup(ctx context.Context, arg CreateActivityGroupParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createActivityGroup, arg.Title, arg.Email)
}

const deleteActivityGroup = `-- name: DeleteActivityGroup :exec
DELETE FROM activity_groups
WHERE id = ?
`

func (q *Queries) DeleteActivityGroup(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteActivityGroup, id)
	return err
}

const getActivityGroup = `-- name: GetActivityGroup :one
SELECT id, title, email, created_at, updated_at FROM activity_groups
WHERE id = ? LIMIT 1
`

func (q *Queries) GetActivityGroup(ctx context.Context, id int32) (ActivityGroup, error) {
	row := q.db.QueryRowContext(ctx, getActivityGroup, id)
	var i ActivityGroup
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listActivityGroups = `-- name: ListActivityGroups :many
SELECT id, title, email, created_at, updated_at FROM activity_groups ORDER BY id DESC
`

func (q *Queries) ListActivityGroups(ctx context.Context) ([]ActivityGroup, error) {
	rows, err := q.db.QueryContext(ctx, listActivityGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ActivityGroup{}
	for rows.Next() {
		var i ActivityGroup
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateActivityGroup = `-- name: UpdateActivityGroup :exec
UPDATE activity_groups
SET title = ?
WHERE id = ?
`

type UpdateActivityGroupParams struct {
	Title string `json:"title"`
	ID    int32  `json:"id"`
}

func (q *Queries) UpdateActivityGroup(ctx context.Context, arg UpdateActivityGroupParams) error {
	_, err := q.db.ExecContext(ctx, updateActivityGroup, arg.Title, arg.ID)
	return err
}
