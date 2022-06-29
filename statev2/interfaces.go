package statev2

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type querier interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}