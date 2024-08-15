package rest

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/cart"
	"github.com/nmarsollier/cartgo/rabbit/emit"
	"github.com/nmarsollier/cartgo/rest/server"
	"github.com/nmarsollier/cartgo/security"
	"github.com/nmarsollier/cartgo/tools/db"
	"github.com/nmarsollier/cartgo/tools/httpx"
	"github.com/stretchr/testify/assert"
)

func TestGetCartCheckoutHappyPath(t *testing.T) {
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

	collection.EXPECT().UpdateOne(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(arg1 interface{}, filter cart.DbIdFilter, update cart.DbDeleteTokenDocument) (int64, error) {
			// Check parameters
			assert.Equal(t, cartData.ID, filter.ID)

			// Asign return values
			return 1, nil
		},
	).Times(1)

	// Security
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpToken(httpMock, user)

	// Service
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("")), // or use an io.NopCloser with a buffer for more control
	}
	httpMock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(2)

	rabbitMock := emit.DefaultRabbitChannel(ctrl, 1)
	rabbitMock.EXPECT().Publish(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(exchange string, routingKey string, body []byte) error {
			assert.Equal(t, "order", exchange)
			assert.Equal(t, "order", routingKey)
			bodyStr := string(body)
			assert.Contains(t, bodyStr, "place-order")
			assert.Contains(t, bodyStr, "cartId")
			assert.Contains(t, bodyStr, "userId")

			return nil
		},
	).Times(1)

	// REQUEST
	r := server.TestRouter(collection, httpMock, rabbitMock)
	InitRoutes()

	req, w := server.TestPostRequest("/v1/cart/checkout", "", user.ID)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	result := w.Body.String()
	assert.NotEmpty(t, result)
}

func TestGetCartCheckoutInvalidToken(t *testing.T) {
	user := security.TestUser()

	// Security
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	security.ExpectHttpUnauthorized(httpMock)

	// REQUEST
	r := server.TestRouter(httpMock)
	InitRoutes()

	req, w := server.TestPostRequest("/v1/cart/checkout", "", user.ID)
	r.ServeHTTP(w, req)

	server.AssertUnauthorized(t, w)
}
