package consume

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/stretchr/testify/assert"
)

func TestProccessArticleDataHappyPath(t *testing.T) {
	cartData := cart.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.DbUserIdFilter, updated *cart.Cart) error {
			assert.Equal(t, 2, len(cartData.Articles))

			*updated = *cartData
			return nil
		},
	).Times(1)

	collection.EXPECT().ReplaceOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, filter cart.DbIdFilter, replaced *cart.Cart) (int64, error) {
			assert.Equal(t, cartData.ID, filter.ID)
			assert.Equal(t, 2, len(replaced.Articles))
			return 1, nil
		},
	).Times(1)

	// REQUEST
	newMessage := &consumeArticleExistMessage{
		Message: &cart.ValidationEvent{
			ReferenceId: "123",
			ArticleId:   "123",
			Valid:       true,
		},
	}

	processArticleExist(newMessage, collection)

}
