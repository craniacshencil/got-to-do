// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: tasks.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (task_id, list_id, task_name, start_time, end_time, completion)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING task_id, list_id, task_name, start_time, end_time, completion
`

type CreateTaskParams struct {
	TaskID     uuid.UUID
	ListID     uuid.UUID
	TaskName   string
	StartTime  time.Time
	EndTime    time.Time
	Completion bool
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.TaskID,
		arg.ListID,
		arg.TaskName,
		arg.StartTime,
		arg.EndTime,
		arg.Completion,
	)
	var i Task
	err := row.Scan(
		&i.TaskID,
		&i.ListID,
		&i.TaskName,
		&i.StartTime,
		&i.EndTime,
		&i.Completion,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :execrows
DELETE from tasks  
WHERE task_id=$1
`

func (q *Queries) DeleteTask(ctx context.Context, taskID uuid.UUID) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteTask, taskID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getTasks = `-- name: GetTasks :many
SELECT task_id, list_id, task_name, start_time, end_time, completion from tasks 
WHERE list_id=$1
`

func (q *Queries) GetTasks(ctx context.Context, listID uuid.UUID) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasks, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.TaskID,
			&i.ListID,
			&i.TaskName,
			&i.StartTime,
			&i.EndTime,
			&i.Completion,
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

const updateTask = `-- name: UpdateTask :exec
UPDATE tasks 
SET 
  task_name=$1, 
  start_time=$2, 
  end_time=$3, 
  completion=$4 
WHERE task_id=$5
`

type UpdateTaskParams struct {
	TaskName   string
	StartTime  time.Time
	EndTime    time.Time
	Completion bool
	TaskID     uuid.UUID
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) error {
	_, err := q.db.ExecContext(ctx, updateTask,
		arg.TaskName,
		arg.StartTime,
		arg.EndTime,
		arg.Completion,
		arg.TaskID,
	)
	return err
}
