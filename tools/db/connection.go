package db

import (
	"context"
	"reflect"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/log"
)

var (
	instance *pgxpool.Pool
)

func GetPostgresClient(deps ...interface{}) (*pgxpool.Pool, error) {
	if instance == nil {
		config, err := pgxpool.ParseConfig(env.Get().PostgresURL)
		if err != nil {
			log.Get(deps...).Error(err)
			return nil, err
		}

		instance, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Get(deps...).Error(err)
			return nil, err
		}

		log.Get(deps...).Info("Postgres Connected")
	}

	return instance, nil
}

func QueryRow[T any](query string, args []interface{}, deps ...interface{}) (result *T, err error) {
	result = new(T)
	destVal := reflect.ValueOf(result).Elem()

	conn, err := GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	row := conn.QueryRow(context.Background(), query, args...)

	if err = row.Scan(fieldAddrs(destVal)...); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	return
}

func fieldAddrs(destVal reflect.Value) (addrs []interface{}) {
	addrs = make([]interface{}, destVal.NumField())
	for i := 0; i < destVal.NumField(); i++ {
		addrs[i] = destVal.Field(i).Addr().Interface()
	}
	return
}
