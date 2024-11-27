package security

import (
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/log"
	gocache "github.com/patrickmn/go-cache"
)

// Validate valida si el token es valido
func Validate(token string, deps ...interface{}) (*User, error) {
	// Si esta en cache, retornamos el cache
	if found, ok := cache.Get(token); ok {
		if user, ok := found.(*User); ok {
			return user, nil
		}
	}

	user, err := getRemoteToken(token, deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, errs.Unauthorized
	}

	// Todo bien, se agrega al cache y se retorna
	cache.Set(token, user, gocache.DefaultExpiration)

	return user, nil
}
