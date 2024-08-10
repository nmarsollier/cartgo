package cart

import (
	"context"
	"time"

	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Define mongo Collection
var collection *mongo.Collection

func dbCollection() (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	col := database.Collection("cart")

	collection = col
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

type findByUserIdFilter struct {
	UserId  string `bson:"userId"`
	Enabled bool   `bson:"enabled"`
}

// findByCartId lee un usuario desde la db
func findByUserId(userId string) (*Cart, error) {
	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	cart := &Cart{}
	filter := findByUserIdFilter{
		UserId:  userId,
		Enabled: true,
	}

	if err = collection.FindOne(context.Background(), filter).Decode(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

// findByCartId lee un usuario desde la db
func findById(cartId string) (*Cart, error) {
	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(cartId)
	if err != nil {
		return nil, errors.ErrID
	}

	cart := &Cart{}
	filter := bson.M{"_id": _id}

	if err = collection.FindOne(context.Background(), filter).Decode(cart); err != nil {
		return nil, err
	}

	return cart, nil
}

func insert(cart *Cart) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		return nil, err
	}

	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	if _, err := collection.InsertOne(context.Background(), cart); err != nil {
		return nil, err
	}

	return cart, nil
}

func replace(cart *Cart) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		return nil, err
	}

	var collection, err = dbCollection()
	if err != nil {
		return nil, err
	}

	_, err = collection.ReplaceOne(context.Background(), bson.M{"_id": cart.ID}, cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func invalidate(cart *Cart) (*Cart, error) {
	if err := cart.ValidateSchema(); err != nil {
		return nil, err
	}

	var collection, err = dbCollection()
	if err != nil {
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
