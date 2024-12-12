package db

import (
	"reflect"

	"github.com/nmarsollier/cartgo/tools/log"
)

func GetQuery[T any](deps ...interface{}) Query[T] {
	for _, o := range deps {
		if ti, ok := o.(Query[T]); ok {
			return ti
		}
	}

	return &pgQuery[T]{
		deps: deps,
	}
}

type Query[T any] interface {
	Row(query string, args ...interface{}) (result *T, err error)
}

type pgQuery[T any] struct {
	deps []interface{}
}

func (q *pgQuery[T]) Row(query string, args ...interface{}) (result *T, err error) {
	result = new(T)
	destVal := reflect.ValueOf(result).Elem()

	conn, err := getDBConn(q.deps...)
	if err != nil {
		log.Get(q.deps...).Error(err)
		return
	}

	row := conn.QueryRow(query, args...)

	if err = row.Scan(fieldAddrs(destVal)...); err != nil {
		if isConnectionError(err) {
			instance = nil
		}

		log.Get(q.deps...).Error(err)
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
