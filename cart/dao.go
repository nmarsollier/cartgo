package cart

import (
	"context"
	"time"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/tools/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define mongo Collection
var collection db.MongoCollection

func dbCollection(ctx ...interface{}) (db.MongoCollection, error) {
	for _, o := range ctx {
		if coll, ok := o.(db.MongoCollection); ok {
			return coll, nil
		}
	}

	if collection != nil {
		return collection, nil
	}

	database, err := db.Get()
	if err != nil {
		glog.Error(err)
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
func findByUserId(userId string, ctx ...interface{}) (*Cart, error) {
	var collection, err = dbCollection(ctx...)
	if err != nil {
		glog.Error(err)
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
func findById(cartId string, ctx ...interface{}) (*Cart, error) {
	_id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		glog.Error(err)
		return nil, ErrID
	}

	collection, err := dbCollection(ctx...)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	cart := &Cart{}
	filter := DbIdFilter{ID: _id}

	if err = collection.FindOne(context.Background(), filter, cart); err != nil {
		glog.Error(err)
		return nil, err
	}

	return cart, nil
}

func insert(cart *Cart, ctx ...interface{}) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		glog.Error(err)
		return nil, err
	}

	var collection, err = dbCollection(ctx...)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	if _, err := collection.InsertOne(context.Background(), cart); err != nil {
		glog.Error(err)
		return nil, err
	}

	return cart, nil
}

func replace(cart *Cart, ctx ...interface{}) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		glog.Error(err)
		return nil, err
	}

	var collection, err = dbCollection(ctx...)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	_, err = collection.ReplaceOne(context.Background(), DbIdFilter{ID: cart.ID}, cart)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return cart, nil
}

func invalidate(cart *Cart, ctx ...interface{}) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		glog.Error(err)
		return nil, err
	}

	var collection, err = dbCollection(ctx...)
	if err != nil {
		glog.Error(err)
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
