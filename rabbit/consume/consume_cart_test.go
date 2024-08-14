package consume

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestProccessArticleDataHappyPath(t *testing.T) {
	cartData := tests.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.FindByUserIdFilter, updated *cart.Cart) error {
			assert.Equal(t, 2, len(cartData.Articles))

			*updated = *cartData
			return nil
		},
	).Times(1)

	collection.EXPECT().ReplaceOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, filter primitive.M, replaced *cart.Cart) (int64, error) {
			assert.Equal(t, cartData.ID, filter["_id"])
			assert.Equal(t, 2, len(replaced.Articles))
			return 1, nil
		},
	).Times(1)

	// REQUEST
	newMessage := &consumeArticleDataMessage{
		Type:     "article-data",
		Queue:    "orders",
		Exchange: "orders",
		Message: &cart.ValidationEvent{
			ReferenceId: "123",
			ArticleId:   "123",
			Valid:       true,
		},
	}

	processArticleData(newMessage, collection)

}
