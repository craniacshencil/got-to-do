// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: todo_lists.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createList = `-- name: CreateList :one
INSERT INTO todo_lists (list_id, user_id, date)
VALUES ($1, $2, $3)
RETURNING list_id, user_id, date
`

type CreateListParams struct {
	ListID uuid.UUID
	UserID uuid.NullUUID
	Date   time.Time
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) (TodoList, error) {
	row := q.db.QueryRowContext(ctx, createList, arg.ListID, arg.UserID, arg.Date)
	var i TodoList
	err := row.Scan(&i.ListID, &i.UserID, &i.Date)
	return i, err
}
