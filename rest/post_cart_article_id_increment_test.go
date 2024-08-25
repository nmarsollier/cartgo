package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/log"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/httpx"
	"github.com/stretchr/testify/assert"
)

func TestPostCartArticleIdIncrementHappyPath2(t *testing.T) {
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
			assert.Equal(t, 3, replaced.Articles[1].Quantity)
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock, log.NewTestLogger(ctrl, 14, 0, 3, 3))
	InitRoutes()

	req, w := server.TestPostRequest("/v1/cart/article/"+cartData.Articles[1].ArticleId+"/increment", "", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestPostCartArticleIdIncrementInvalidToken(t *testing.T) {
	user := security.TestUser()
	cartData := cart.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock, log.NewTestLogger(ctrl, 5, 2, 1, 1))
	InitRoutes()

	req, w := server.TestPostRequest("/v1/cart/article/"+cartData.Articles[1].ArticleId+"/increment", "", user.ID)
	r.ServeHTTP(w, req)

	server.AssertUnauthorized(t, w)
}

func TestPostCartArticleIdIncrementDocumentNotFound(t *testing.T) {
	user := security.TestUser()
	cartData := cart.TestCart()

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

	req, w := server.TestPostRequest("/v1/cart/article/"+cartData.Articles[0].ArticleId+"/increment", "", user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}

// Test que elimia articulo con stock 0
func TestPostCartArticleIdIncrementReplaceError(t *testing.T) {
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

	req, w := server.TestPostRequest("/v1/cart/article/"+cartData.Articles[0].ArticleId+"/increment", "", user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}
