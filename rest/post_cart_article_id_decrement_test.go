package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/http_client"
	"github.com/nmarsollier/cartgo/tools/tests"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test que elimia articulo con stock 0
func TestPostCartArticleIdDecrementHappyPath1(t *testing.T) {
	user := security.TestUser()
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
			assert.Equal(t, 1, len(replaced.Articles))
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/cart/article/"+cartData.Articles[0].ArticleId+"/decrement", "", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostCartArticleIdDecrementHappyPath2(t *testing.T) {
	user := security.TestUser()
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
			assert.Equal(t, 1, replaced.Articles[1].Quantity)
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/cart/article/"+cartData.Articles[1].ArticleId+"/decrement", "", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostCartArticleIdDecrementInvalidToken(t *testing.T) {
	user := security.TestUser()
	cartData := tests.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/cart/article/"+cartData.Articles[1].ArticleId+"/decrement", "", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertUnauthorized(t, w)
}

func TestPostCartArticleIdDecrementDocumentNotFound(t *testing.T) {
	user := security.TestUser()
	cartData := tests.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	tests.ExpectFindOneError(collection, errs.NotFound, 1)

	// Security
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/cart/article/"+cartData.Articles[0].ArticleId+"/decrement", "", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}

// Test que elimia articulo con stock 0
func TestPostCartArticleIdDecrementReplaceError(t *testing.T) {
	user := security.TestUser()
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

	tests.ExpectReplaceOneError(collection, errs.NotFound, 1)

	// Security
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestPostRequest("/v1/cart/article/"+cartData.Articles[0].ArticleId+"/decrement", "", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}
