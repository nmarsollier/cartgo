package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rabbit/emit"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/httpx"
	"github.com/nmarsollier/cartgo/tools/log"
	"github.com/stretchr/testify/assert"
)

func TestPostCartArticleHappyPath(t *testing.T) {
	user := security.TestUser()
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
			assert.Equal(t, 3, len(replaced.Articles))
			assert.Equal(t, 1, replaced.Articles[0].Quantity)
			assert.Equal(t, 2, replaced.Articles[1].Quantity)
			assert.Equal(t, 10, replaced.Articles[2].Quantity)
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	rabbitMock := emit.DefaultRabbitChannel(ctrl, 3)
	rabbitMock.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(exchange string, routingKey string, body []byte) error {
			assert.Equal(t, "article_exist", exchange)
			return nil
		},
	).Times(3)

	// REQUEST
	r := server.TestRouter(collection, httpMock, rabbitMock, log.NewTestLogger(ctrl, 18, 0, 4, 4))
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: "new_1",
		Quantity:  10,
	}
	req, w := server.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostCartArticleHappyPath2(t *testing.T) {
	user := security.TestUser()
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
			assert.Equal(t, 11, replaced.Articles[0].Quantity)
			assert.Equal(t, 2, replaced.Articles[1].Quantity)
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	rabbitMock := emit.DefaultRabbitChannel(ctrl, 2)
	rabbitMock.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(exchange string, routingKey string, body []byte) error {
			assert.Equal(t, "article_exist", exchange)
			return nil
		},
	).Times(2)

	// REQUEST
	r := server.TestRouter(collection, httpMock, rabbitMock, log.NewTestLogger(ctrl, 14, 0, 3, 3))
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: cartData.Articles[0].ArticleId,
		Quantity:  10,
	}
	req, w := server.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostCartArticleInvalidToken(t *testing.T) {
	user := security.TestUser()
	cartData := cart.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock, log.NewTestLogger(ctrl, 5, 2, 1, 1))
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: cartData.Articles[0].ArticleId,
		Quantity:  10,
	}
	req, w := server.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	server.AssertUnauthorized(t, w)
}

func TestPostCartArticleDocumentNotFound(t *testing.T) {
	user := security.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	db.ExpectFindOneError(collection, errs.NotFound, 1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock, log.NewTestLogger(ctrl, 6, 0, 1, 1))
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: "new_1",
		Quantity:  10,
	}
	req, w := server.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}

func TestPostCartArticleReplaceError(t *testing.T) {
	user := security.TestUser()
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

	db.ExpectReplaceOneError(collection, errs.NotFound, 1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock, log.NewTestLogger(ctrl, 6, 1, 1, 1))
	InitRoutes()

	body := cart.AddArticleData{
		ArticleId: "new_1",
		Quantity:  10,
	}
	req, w := server.TestPostRequest("/v1/cart/article", body, user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}
