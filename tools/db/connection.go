package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/log"
)

var (
	instance DBConn
)

type DBConn interface {
	QueryRow(sql string, args ...any) pgx.Row
	Exec(query string, args ...interface{}) (err error)
	Reset()
}

func getDBConn(deps ...interface{}) (DBConn, error) {
	for _, o := range deps {
		if ti, ok := o.(DBConn); ok {
			return ti, nil
		}
	}

	if instance != nil {
		return instance, nil
	}

	pool, err := postgresConn(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	instance = &pgConn{
		conn: pool,
	}

	return instance, nil
}

type pgConn struct {
	conn *pgxpool.Pool
}

func (c *pgConn) QueryRow(sql string, args ...any) pgx.Row {
	return c.conn.QueryRow(context.Background(), sql, args...)
}

func (c *pgConn) Exec(query string, arguments ...interface{}) (err error) {
	_, err = c.conn.Exec(context.Background(), query, arguments...)
	return
}

func (c *pgConn) Reset() {
	instance = nil
}

func postgresConn(deps ...interface{}) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(env.Get().PostgresURL)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	log.Get(deps...).Info("Postgres Connected")

	return pool, nil
}

func isConnectionError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "08000", "08003", "08006", "08001", "08004", "08007", "08P01":
			return true
		}
	}

	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return true
	}

	return false
}
