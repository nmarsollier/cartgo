package db

import (
	"github.com/nmarsollier/cartgo/tools/log"
)

func GetUpdate(deps ...interface{}) Update {
	for _, o := range deps {
		if ti, ok := o.(Update); ok {
			return ti
		}
	}

	return &pgUpdate{
		deps: deps,
	}
}

type Update interface {
	Exec(query string, args ...interface{}) (err error)
}

type pgUpdate struct {
	deps []interface{}
}

func (u *pgUpdate) Exec(query string, arguments ...any) (err error) {
	conn, err := getDBConn(u.deps...)
	if err != nil {
		log.Get(u.deps...).Error(err)
		return err
	}

	err = conn.Exec(query, arguments...)
	if err != nil {
		if isConnectionError(err) {
			instance = nil
		}

		log.Get(u.deps...).Error(err)
		return err
	}

	return nil
}
