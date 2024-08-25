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
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetUsersHappyPath(t *testing.T) {
	user := security.TestUser()
	cartData := cart.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.DbUserIdFilter, updated *cart.Cart) error {
			// Check parameters
			assert.Equal(t, user.ID, params.UserId)
			assert.Equal(t, true, params.Enabled)

			// Asign return values
			*updated = *cartData
			return nil
		},
	).Times(1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock, log.NewTestLogger(ctrl, 14, 0, 3, 3))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetUsersNewCartHappyPath(t *testing.T) {
	user := security.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.DbUserIdFilter, updated *cart.Cart) error {
			return mongo.ErrNoDocuments
		},
	).Times(1)

	collection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, updated *cart.Cart) (interface{}, error) {
			// Check parameters
			assert.Equal(t, true, updated.Enabled)
			assert.Equal(t, "", updated.OrderId)
			return "", nil
		},
	).Times(1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock, log.NewTestLogger(ctrl, 6, 0, 1, 1))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetUsersInsertDbError(t *testing.T) {
	user := security.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.DbUserIdFilter, updated *cart.Cart) error {
			return mongo.ErrNoDocuments
		},
	).Times(1)

	db.ExpectInsertOneError(collection, errs.Internal, 1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// REQUEST
	r := server.TestRouter(collection, httpMock, log.NewTestLogger(ctrl, 6, 1, 1, 1))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	server.AssertInternalServerError(t, w)
}

func TestGetUsersDbError(t *testing.T) {
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

	req, w := server.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	server.AssertDocumentNotFound(t, w)
}

func TestGetUsersTokenInvalid(t *testing.T) {
	user := security.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock, log.NewTestLogger(ctrl, 5, 2, 1, 1))
	InitRoutes()

	req, w := server.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	server.AssertUnauthorized(t, w)
}
