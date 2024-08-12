package rest

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rest/engine"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/service"
	"github.com/nmarsollier/cartgo/tools/apperr"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/tests"
	"github.com/stretchr/testify/assert"
)

func TestGetCartValidateHappyPath(t *testing.T) {
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

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// Serice
	serviceMock := service.NewMockServiceDao(ctrl)
	serviceMock.EXPECT().CallValidate(gomock.Any(), gomock.Any()).Return(nil).Times(2)

	// REQUEST
	r := engine.TestRouter(cart.NewCartOptions(collection), httpMock, serviceMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetCartValidateDocumentNotFound(t *testing.T) {
	user := tests.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	collection := db.NewMockMongoCollection(ctrl)
	tests.ExpectFindOneError(collection, apperr.NotFound, 1)

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// REQUEST
	r := engine.TestRouter(cart.NewCartOptions(collection), httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertDocumentNotFound(t, w)
}

func TestGetCartValidateInvalidToken(t *testing.T) {
	user := tests.TestUser()

	// DB Mock
	ctrl := gomock.NewController(t)
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(nil, apperr.Unauthorized)

	// REQUEST
	r := engine.TestRouter(httpMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertUnauthorized(t, w)
}

func TestGetCartValidateInvalidArticleAth(t *testing.T) {
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

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// Serice
	serviceMock := service.NewMockServiceDao(ctrl)
	serviceMock.EXPECT().CallValidate(gomock.Any(), gomock.Any()).Return(apperr.Unauthorized).Times(1)

	// REQUEST
	r := engine.TestRouter(cart.NewCartOptions(collection), httpMock, serviceMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertUnauthorized(t, w)
}

func TestGetCartValidateInvalidArticle(t *testing.T) {
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

	// Security
	httpMock := security.NewMockSecurityDao(ctrl)
	httpMock.EXPECT().GetRemoteToken(gomock.Any()).Return(user, nil)

	// Serice
	serviceMock := service.NewMockServiceDao(ctrl)
	serviceMock.EXPECT().CallValidate(gomock.Any(), gomock.Any()).Return(apperr.Invalid).Times(1)

	// REQUEST
	r := engine.TestRouter(cart.NewCartOptions(collection), httpMock, serviceMock)
	InitRoutes()

	req, w := tests.TestGetRequest("/v1/cart/validate", user.ID)
	r.ServeHTTP(w, req)

	tests.AssertBadRequestError(t, w)
}
