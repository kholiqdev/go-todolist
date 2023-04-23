// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package querier

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateActivityGroup(ctx context.Context, arg CreateActivityGroupParams) (sql.Result, error)
	DeleteActivityGroup(ctx context.Context, id int32) error
	GetActivityGroup(ctx context.Context, id int32) (Activity, error)
	ListActivityGroups(ctx context.Context) ([]Activity, error)
	UpdateActivityGroup(ctx context.Context, arg UpdateActivityGroupParams) error
}

var _ Querier = (*Queries)(nil)
