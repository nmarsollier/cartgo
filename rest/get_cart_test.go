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
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetUsersHappyPath(t *testing.T) {
	user := tests.TestUser()
	cartData := tests.TestCart()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.FindByUserIdFilter, updated *cart.Cart) error {
			// Check parameters
			assert.Equal(t, user.ID, params.UserId)
			assert.Equal(t, true, params.Enabled)

			// Asign return values
			*updated = *cartData
			return nil
		},
	).Times(1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetUsersNewCartHappyPath(t *testing.T) {
	user := tests.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.FindByUserIdFilter, updated *cart.Cart) error {
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
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetUsersInsertDbError(t *testing.T) {
	user := tests.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	collection.EXPECT().FindOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, params cart.FindByUserIdFilter, updated *cart.Cart) error {
			return mongo.ErrNoDocuments
		},
	).Times(1)

	tests.ExpectInsertOneError(collection, apperr.Internal, 1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertInternalServerError(t, w)
}

func TestGetUsersDbError(t *testing.T) {
	user := tests.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	tests.ExpectFindOneError(collection, apperr.NotFound, 1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}

func TestGetUsersTokenInvalid(t *testing.T) {
	user := tests.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(nil, apperr.Unauthorized)

	// REQUEST
	r := engine.TestRouter(httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertUnauthorized(t, w)
}
