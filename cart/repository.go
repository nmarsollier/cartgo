package cart

import (
	"context"
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
	query := `
        SELECT id, userId, orderId, articles, enabled, created, updated
        FROM cartgo.carts
        WHERE userId = $1 and enabled = true
    `

	cart, err := db.QueryRow[Cart](query, []interface{}{userId}, deps...)
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
	query := `
        SELECT id, userId, orderId, articles, enabled, created, updated
        FROM cartgo.carts
        WHERE id = $1
    `
	cart, err := db.QueryRow[Cart](query, []interface{}{cartId}, deps...)
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

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	articlesJSON := strs.ToJson(cart.Articles)

	query := `
        INSERT INTO cartgo.carts (id, userId, orderId, articles, enabled, created, updated)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (id) DO UPDATE SET
            articles = EXCLUDED.articles,
            enabled = EXCLUDED.enabled,
            updated = EXCLUDED.updated
    `

	_, err = conn.Exec(context.Background(), query, cart.ID, cart.UserId, cart.OrderId, articlesJSON, cart.Enabled, cart.Created, cart.Updated)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	return nil
}

func invalidate(cart *Cart, deps ...interface{}) (err error) {
	if err = cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	conn, err := db.GetPostgresClient(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	query := `
        UPDATE cartgo.carts
        SET enabled = $1, updated = $2
        WHERE id = $3
    `

	_, err = conn.Exec(context.Background(), query, false, time.Now(), cart.ID)
	if err != nil {
		log.Get(deps...).Error(err)
		return err
	}

	return nil
}
