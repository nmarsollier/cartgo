package cart

import (
	"context"
	"time"

	"github.com/golang/glog"
	"github.com/nmarsollier/cartgo/tools/apperr"
	"github.com/nmarsollier/cartgo/tools/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define mongo Collection
var collection db.MongoCollection

func NewCartOptions(collection db.MongoCollection) CartOption {
	return CartOption{
		Collection: collection,
	}
}

type CartOption struct {
	Collection db.MongoCollection
}

func dbCollection(options ...interface{}) (db.MongoCollection, error) {
	for _, o := range options {
		if ti, ok := o.(CartOption); ok {
			return ti.Collection, nil
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

type FindByUserIdFilter struct {
	UserId  string `bson:"userId"`
	Enabled bool   `bson:"enabled"`
}

// findByCartId lee un usuario desde la db
func findByUserId(userId string, options ...interface{}) (*Cart, error) {
	var collection, err = dbCollection(options...)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	cart := &Cart{}
	filter := FindByUserIdFilter{
		UserId:  userId,
		Enabled: true,
	}

	if err = collection.FindOne(context.Background(), filter, cart); err != nil {
		return nil, err
	}

	return cart, nil
}

// findByCartId lee un usuario desde la db
func findById(cartId string, options ...interface{}) (*Cart, error) {
	_id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		glog.Error(err)
		return nil, apperr.ErrID
	}

	collection, err := dbCollection(options...)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	cart := &Cart{}
	filter := bson.M{"_id": _id}

	if err = collection.FindOne(context.Background(), filter, cart); err != nil {
		glog.Error(err)
		return nil, err
	}

	return cart, nil
}

func insert(cart *Cart, options ...interface{}) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		glog.Error(err)
		return nil, err
	}

	var collection, err = dbCollection(options...)
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

func replace(cart *Cart, options ...interface{}) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		glog.Error(err)
		return nil, err
	}

	var collection, err = dbCollection(options...)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	_, err = collection.ReplaceOne(context.Background(), bson.M{"_id": cart.ID}, cart)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	return cart, nil
}

func invalidate(cart *Cart, options ...interface{}) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		glog.Error(err)
		return nil, err
	}

	var collection, err = dbCollection(options...)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	cart.Enabled = false

	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": cart.ID},
		bson.M{"$set": bson.M{
			"enabled": false,
		}},
	)

	return cart, err
}
