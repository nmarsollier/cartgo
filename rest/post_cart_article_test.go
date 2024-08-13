package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/apperr"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPostCartArticleHappyPath(t *testing.T) {
	user := tests.TestUser()
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
			assert.Equal(t, 3, len(replaced.Articles))
			assert.Equal(t, 1, replaced.Articles[0].Quantity)
			assert.Equal(t, 2, replaced.Articles[1].Quantity)
			assert.Equal(t, 10, replaced.Articles[2].Quantity)
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(cart.CartCollection(collection), httpMock)
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: "new_1",
		Quantity:  10,
	}
	req, w := tests.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostCartArticleHappyPath2(t *testing.T) {
	user := tests.TestUser()
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
			assert.Equal(t, 11, replaced.Articles[0].Quantity)
			assert.Equal(t, 2, replaced.Articles[1].Quantity)
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(cart.CartCollection(collection), httpMock)
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: cartData.Articles[0].ArticleId,
		Quantity:  10,
	}
	req, w := tests.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostCartArticleInvalidToken(t *testing.T) {
	user := tests.TestUser()
	cartData := tests.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(nil, apperr.Unauthorized)

	// REQUEST
	r := engine.TestRouter(httpMock)
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: cartData.Articles[0].ArticleId,
		Quantity:  10,
	}
	req, w := tests.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	tests.AssertUnauthorized(t, w)
}

func TestPostCartArticleDocumentNotFound(t *testing.T) {
	user := tests.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	tests.ExpectFindOneError(collection, apperr.NotFound, 1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(cart.CartCollection(collection), httpMock)
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: "new_1",
		Quantity:  10,
	}
	req, w := tests.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}

func TestPostCartArticleReplaceError(t *testing.T) {
	user := tests.TestUser()
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

	tests.ExpectReplaceOneError(collection, apperr.NotFound, 1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(cart.CartCollection(collection), httpMock)
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: "new_1",
		Quantity:  10,
	}
	req, w := tests.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}