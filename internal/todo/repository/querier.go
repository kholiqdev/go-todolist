// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package querier

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateTodo(ctx context.Context, title string) (sql.Result, error)
	DeleteTodo(ctx context.Context, id int32) error
	GetTodo(ctx context.Context, id int32) (Todo, error)
	ListTodos(ctx context.Context, arg ListTodosParams) ([]Todo, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) error
}

var _ Querier = (*Queries)(nil)
