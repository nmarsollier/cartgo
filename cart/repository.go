package cart

import (
	"context"
	"time"

	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrID = errs.NewValidation().Add("id", "Invalid")

// Define mongo Collection
var collection db.MongoCollection

func dbCollection(deps ...interface{}) (db.MongoCollection, error) {
	for _, o := range deps {
		if coll, ok := o.(db.MongoCollection); ok {
			return coll, nil
		}
	}

	if collection != nil {
		return collection, nil
	}

	database, err := db.Get()
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	col := database.Collection("cart")

	collection = db.NewMongoCollection(col)
	return collection, nil
}

func newCart(userId string) *Cart {
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
func findByUserId(userId string, deps ...interface{}) (*Cart, error) {
	var collection, err = dbCollection(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	cart := &Cart{}
	filter := DbUserIdFilter{
		UserId:  userId,
		Enabled: true,
	}

	if err = collection.FindOne(context.Background(), filter, cart); err != nil {
		return nil, err
	}

	return cart, nil
}

// findByCartId lee un usuario desde la db
func findById(cartId string, deps ...interface{}) (*Cart, error) {
	_id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, ErrID
	}

	collection, err := dbCollection(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	cart := &Cart{}
	filter := DbIdFilter{ID: _id}

	if err = collection.FindOne(context.Background(), filter, cart); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return cart, nil
}

func insert(cart *Cart, deps ...interface{}) (*Cart, error) {
	if err := cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	var collection, err = dbCollection(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	if _, err := collection.InsertOne(context.Background(), cart); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return cart, nil
}

func replace(cart *Cart, deps ...interface{}) (*Cart, error) {
	if err := cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	var collection, err = dbCollection(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	_, err = collection.ReplaceOne(context.Background(), DbIdFilter{ID: cart.ID}, cart)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	return cart, nil
}

func invalidate(cart *Cart, deps ...interface{}) (*Cart, error) {
	if err := cart.validateSchema(); err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	var collection, err = dbCollection(deps...)
	if err != nil {
		log.Get(deps...).Error(err)
		return nil, err
	}

	cart.Enabled = false

	_, err = collection.UpdateOne(context.Background(),
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
