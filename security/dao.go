package security

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nmarsollier/cartgo/log"
	"github.com/nmarsollier/cartgo/tools/env"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/httpx"
)

func getRemoteToken(token string, ctx ...interface{}) (*User, error) {
	// Buscamos el usuario remoto
	req, err := http.NewRequest("GET", env.Get().SecurityServerURL+"/v1/users/current", nil)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, errs.Unauthorized
	}
	req.Header.Add("Authorization", "bearer "+token)
	resp, err := httpx.Get(ctx...).Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Get(ctx...).Error(err)
		return nil, errs.Unauthorized
	}
	defer resp.Body.Close()

	user := &User{}
	err = json.NewDecoder(resp.Body).Decode(user)
	if err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Get(ctx...).Error(err)
		return nil, err
	}
	return user, nil
}
