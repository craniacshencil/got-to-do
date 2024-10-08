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
	UserID uuid.UUID
	Date   time.Time
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) (TodoList, error) {
	row := q.db.QueryRowContext(ctx, createList, arg.ListID, arg.UserID, arg.Date)
	var i TodoList
	err := row.Scan(&i.ListID, &i.UserID, &i.Date)
	return i, err
}

const deleteList = `-- name: DeleteList :execrows
DELETE FROM todo_lists
WHERE list_id=$1
`

func (q *Queries) DeleteList(ctx context.Context, listID uuid.UUID) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteList, listID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getListID = `-- name: GetListID :one
SELECT list_id from todo_lists
WHERE date=$1 and user_id=$2
`

type GetListIDParams struct {
	Date   time.Time
	UserID uuid.UUID
}

func (q *Queries) GetListID(ctx context.Context, arg GetListIDParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getListID, arg.Date, arg.UserID)
	var list_id uuid.UUID
	err := row.Scan(&list_id)
	return list_id, err
}
