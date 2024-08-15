package rest

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/http_client"
	"github.com/nmarsollier/cartgo/tools/str_tools"
	"github.com/nmarsollier/cartgo/tools/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetCartValidateHappyPath(t *testing.T) {
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

	// Security
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// Service
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(str_tools.ToJson(user))),
	}
	httpMock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(2)

	// REQUEST
	r := server.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetCartValidateDocumentNotFound(t *testing.T) {
	user := security.TestUser()

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

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}

func TestGetCartValidateInvalidToken(t *testing.T) {
	user := security.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertUnauthorized(t, w)
}

func TestGetCartValidateInvalidArticleAth(t *testing.T) {
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

	// Security
	httpMock := http_client.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// Service
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(collection, httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertBadRequestError(t, w)
}
