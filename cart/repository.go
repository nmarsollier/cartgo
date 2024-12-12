package cart

import (
	"database/sql"
	"errors"
	"time"

	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/log"
	"github.com/nmarsollier/cartgo/tools/strs"
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
func findByUserId(userId string, deps ...interface{}) (*Cart, error) {
	cart, err := db.GetQuery[Cart](deps...).Row(
		`
      SELECT id, userId, orderId, articles, enabled, created, updated
      FROM cartgo.carts
      WHERE userId = $1 and enabled = true
    `,
		userId,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errs.NotFound
	}

	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return cart, nil
}

func findById(cartId string, deps ...interface{}) (*Cart, error) {
	cart, err := db.GetQuery[Cart](deps...).Row(
		`
      SELECT id, userId, orderId, articles, enabled, created, updated
      FROM cartgo.carts
      WHERE id = $1
    `,
		cartId,
	)

	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return cart, nil
}

func save(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	err = db.GetUpdate(deps...).Exec(
		`
      INSERT INTO cartgo.carts (id, userId, orderId, articles, enabled, created, updated)
      VALUES ($1, $2, $3, $4, $5, $6, $7)
      ON CONFLICT (id) DO UPDATE SET
          articles = EXCLUDED.articles,
          enabled = EXCLUDED.enabled,
          updated = EXCLUDED.updated
    `,
		cart.ID,
		cart.UserId,
		cart.OrderId,
		strs.ToJson(cart.Articles),
		cart.Enabled,
		cart.Created,
		cart.Updated,
	)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}

func invalidate(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return
	}

	err = db.GetUpdate(deps...).Exec(
		`
        UPDATE cartgo.carts
        SET enabled = $1, updated = $2
        WHERE id = $3
    `,
		false,
		time.Now(),
		cart.ID,
	)

	if err != nil {
		log.Get(deps...).Error(err)
	}

	return
}
