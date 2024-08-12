package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/tools/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// Mock Data
var TestIsUniqueError = mongo.WriteException{
	WriteErrors: []mongo.WriteError{
		{
			Code: 11000,
		},
	},
}

func ExpectFindOneError(collection *db.MockMongoCollection, err error, times int) {
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params interface{}, update interface{}) error {
			return err
		},
	).Times(times)
}

func ExpectReplaceOneError(collection *db.MockMongoCollection, err error, times int) {
	collection.EXPECT().ReplaceOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), err).Times(times)
}

func ExpectInsertOneError(collection *db.MockMongoCollection, err error, times int) {
	collection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("", err).Times(times)
}

func ExpectUserInsertOne(collection *db.MockMongoCollection, times int) {
	collection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("123", nil).Times(times)
}

/*
// Espect common functions
func ExpectInsertOneError(userCollection *db.MockMongoCollection, err error, times int) {
	userCollection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("", err).Times(times)
}

func ExpectUserFindOne(userCollection *db.MockMongoCollection, userData *user.User, times int) {
	userCollection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params primitive.M, update *user.User) error {
			*update = *userData
			return nil
		},
	).Times(times)
}
func ExpectUserInsertOne(userCollection *db.MockMongoCollection, times int) {
	userCollection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("123", nil).Times(times)
}

func ExpectUpdateOneError(userCollection *db.MockMongoCollection, err error, times int) {
	userCollection.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), err).Times(times)
}

func ExpectTokenFinOne(tokenCollection *db.MockMongoCollection, tokenData *token.Token, times int) {
	tokenCollection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params primitive.M, token *token.Token) error {
			// Asign return values
			*token = *tokenData
			return nil
		},
	).Times(times)
}

func ExpectTokenInsertOne(tokenCollection *db.MockMongoCollection, times int) {
	tokenCollection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("", nil).Times(times)
}

func ExpectFindOneForToken(t *testing.T, tokenCollection *db.MockMongoCollection, tokenData *token.Token) {
	tokenCollection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params primitive.M, token *token.Token) error {
			assert.Equal(t, tokenData.ID, params["_id"])

			*token = *tokenData
			return nil
		},
	).Times(1)
}
*/
