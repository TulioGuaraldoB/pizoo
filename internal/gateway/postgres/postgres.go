package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/TulioGuaraldoB/pizoo/internal/config/env"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDbPool interface {
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	AcquireAllIdle(ctx context.Context) []*pgxpool.Conn
	AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	Close()
	Config() *pgxpool.Config
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Ping(ctx context.Context) error
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Reset()
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Stat() *pgxpool.Stat
}

var db *pgxpool.Pool

func dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		env.Env.PostgresUser,
		env.Env.PostgresPassword,
		env.Env.PostgresHost,
		env.Env.PostgresPort,
		env.Env.PostgresDatabase,
	)
}

func startDatabase() {
	database, err := pgxpool.New(context.Background(), dsn())
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL Server. %s", err.Error())
	}

	db = database
}

func OpenConnection() IDbPool {
	if db == nil {
		startDatabase()
	}

	return db
}
