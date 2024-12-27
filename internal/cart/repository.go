package cart

import (
	"context"
	"time"

	"github.com/nmarsollier/cartgo/internal/engine/db"
	"github.com/nmarsollier/cartgo/internal/engine/errs"
	"github.com/nmarsollier/cartgo/internal/engine/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartRepository interface {
	FindByUserId(userId string) (*Cart, error)
	FindById(cartId string) (*Cart, error)
	Insert(crt *Cart) (*Cart, error)
	Replace(crt *Cart) (*Cart, error)
	Invalidate(crt *Cart) (*Cart, error)
	NewCart(userId string) *Cart
}

func NewCartRepository(log log.LogRusEntry, collection db.Collection) CartRepository {
	return &cartRepository{
		log:        log,
		collection: collection,
	}
}

type cartRepository struct {
	log        log.LogRusEntry
	collection db.Collection
}

var ErrID = errs.NewValidation().Add("id", "Invalid")

func (r *cartRepository) NewCart(userId string) *Cart {
	return &Cart{
		ID:       primitive.NewObjectID(),
		UserId:   userId,
		Enabled:  true,
		Created:  time.Now(),
		Updated:  time.Now(),
		Articles: []*Article{},
	}
}

type DbUserIdFilter struct {
	UserId  string `bson:"userId"`
	Enabled bool   `bson:"enabled"`
}

type DbIdFilter struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
}

// findByCartId lee un usuario desde la db
func (r *cartRepository) FindByUserId(userId string) (*Cart, error) {
	cart := &Cart{}
	filter := DbUserIdFilter{
		UserId:  userId,
		Enabled: true,
	}

	if err := r.collection.FindOne(context.Background(), filter, cart); err != nil {
		return nil, err
	}

	return cart, nil
}

// findByCartId lee un usuario desde la db
func (r *cartRepository) FindById(cartId string) (*Cart, error) {
	_id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		r.log.Error(err)
		return nil, ErrID
	}

	cart := &Cart{}
	filter := DbIdFilter{ID: _id}

	if err = r.collection.FindOne(context.Background(), filter, cart); err != nil {
		r.log.Error(err)
		return nil, err
	}

	return cart, nil
}

func (r *cartRepository) Insert(cart *Cart) (*Cart, error) {
	if err := cart.validateSchema(); err != nil {
		r.log.Error(err)
		return nil, err
	}

	if _, err := r.collection.InsertOne(context.Background(), cart); err != nil {
		r.log.Error(err)
		return nil, err
	}

	return cart, nil
}

func (r *cartRepository) Replace(cart *Cart) (*Cart, error) {
	if err := cart.validateSchema(); err != nil {
		r.log.Error(err)
		return nil, err
	}

	_, err := r.collection.ReplaceOne(context.Background(), DbIdFilter{ID: cart.ID}, cart)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}

	return cart, nil
}

func (r *cartRepository) Invalidate(cart *Cart) (*Cart, error) {
	if err := cart.validateSchema(); err != nil {
		r.log.Error(err)
		return nil, err
	}

	cart.Enabled = false

	_, err := r.collection.UpdateOne(context.Background(),
		DbIdFilter{ID: cart.ID},
		DbDeleteTokenDocument{
			DbDeleteTokenBody{
				Enabled: false,
			},
		},
	)

	return cart, err
}

type DbDeleteTokenBody struct {
	Enabled bool `bson:"enabled"`
}

type DbDeleteTokenDocument struct {
	Set DbDeleteTokenBody `bson:"$set"`
}
