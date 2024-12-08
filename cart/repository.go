package cart

import (
	"time"

	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/log"
	uuid "github.com/satori/go.uuid"
)

var ErrID = errs.NewValidation().Add("id", "Invalid")

func newCart(userId string) *Cart {
	return &Cart{
		ID:       uuid.NewV4().String(),
		UserId:   userId,
		Enabled:  true,
		Created:  time.Now(),
		Updated:  time.Now(),
		Articles: []*Article{},
	}
}

// findByUserId lee el cart activo del usuario
func findByUserId(cartId string, deps ...interface{}) (cart *Cart, err error) {
	conn, err := GetCartDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	if cart, err = conn.FindByUserId(cartId); err != nil {
		return nil, err
	}

	return cart, nil
}

func findById(cartId string, deps ...interface{}) (cart *Cart, err error) {
	conn, err := GetCartDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	if cart, err = conn.FindById(cartId); err != nil {
		return nil, err
	}

	return cart, nil
}

func insert(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	collection, err := GetCartDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	if err = collection.Save(cart); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	return
}

func replace(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	collection, err := GetCartDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	if err = collection.Save(cart); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	return
}

func invalidate(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	conn, err := GetCartDao(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return
	}

	cart.Enabled = false

	err = conn.Disable(cart.ID, cart.UserId)
	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}
