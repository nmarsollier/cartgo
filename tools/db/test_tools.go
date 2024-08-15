package db

import (
	"github.com/golang/mock/gomock"
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

func ExpectFindOneError(collection *MockMongoCollection, err error, times int) {
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params interface{}, update interface{}) error {
			return err
		},
	).Times(times)
}

func ExpectReplaceOneError(collection *MockMongoCollection, err error, times int) {
	collection.EXPECT().ReplaceOne(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(0), err).Times(times)
}

func ExpectInsertOneError(collection *MockMongoCollection, err error, times int) {
	collection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("", err).Times(times)
}

func ExpectUserInsertOne(collection *MockMongoCollection, times int) {
	collection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return("123", nil).Times(times)
}
