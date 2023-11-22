package postgres

import (
	"context"

	// db "exchange-rate-api/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DBPG interface use function from pgx is so that make 'NewPgPool()' as
// method of 'PgxPool' struct
type DBPG interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// struct that hold the pool connection
type DB struct {
	// db DBPG
	db *pgxpool.Pool
}

// parsing the pool connection to struct
func NewDB(db *pgxpool.Pool) *DB {
	return &DB{
		db: db,
	}
}
