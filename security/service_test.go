package security

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nmarsollier/cartgo/log"
	"github.com/nmarsollier/cartgo/tools/errs"
	"github.com/nmarsollier/cartgo/tools/httpx"
	"gopkg.in/go-playground/assert.v1"
)

func TestInvalidateHappyPath(t *testing.T) {
	testUser := TestUser()
	token := "bearer " + testUser.ID

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	ExpectHttpToken(httpMock, testUser)
	ExpectHttpToken(httpMock, testUser)

	// REQUEST
	logm := log.NewTestLogger(ctrl, 0, 0, 1, 2)
	user, err := Validate(token, httpMock, logm)
	assert.Equal(t, testUser.ID, user.ID)
	assert.Equal(t, nil, err)
	Invalidate(token, logm)
	user, err = Validate(token, httpMock, logm)
	assert.Equal(t, testUser.ID, user.ID)
	assert.Equal(t, nil, err)
}

func TestInvalidateNotAuthorized(t *testing.T) {
	testUser := TestUser()
	token := "bearer " + testUser.ID

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	ExpectHttpToken(httpMock, testUser)
	ExpectHttpUnauthorized(httpMock)

	// REQUEST
	logm := log.NewTestLogger(ctrl, 0, 2, 1, 2)
	user, err := Validate(token, httpMock, logm)
	assert.Equal(t, testUser.ID, user.ID)
	assert.Equal(t, nil, err)
	Invalidate(token, logm)
	user, err = Validate(token, httpMock, logm)
	assert.Equal(t, nil, user)
	assert.Equal(t, errs.Unauthorized, err)
}

func TestInvalidateNotAuthorized2(t *testing.T) {
	testUser := TestUser()
	token := "bearer " + testUser.ID

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	ExpectHttpUnauthorized(httpMock)
	ExpectHttpToken(httpMock, testUser)

	// REQUEST
	logm := log.NewTestLogger(ctrl, 0, 2, 1, 2)
	user, err := Validate(token, httpMock, logm)
	assert.Equal(t, nil, user)
	assert.Equal(t, errs.Unauthorized, err)
	Invalidate(token, logm)
	user, err = Validate(token, httpMock, logm)
	assert.Equal(t, testUser.ID, user.ID)
	assert.Equal(t, nil, err)
}

func TestInvalidateInvalidData(t *testing.T) {
	testUser := TestUser()
	token := "bearer " + testUser.ID

	// Mocks
	ctrl := gomock.NewController(t)
	httpMock := httpx.NewMockHTTPClient(ctrl)
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("123")),
	}
	httpMock.EXPECT().Do(gomock.Any()).Return(response, nil).Times(1)

	// REQUEST
	user, err := Validate(token, httpMock, log.NewTestLogger(ctrl, 0, 2, 0, 1))
	assert.Equal(t, nil, user)
	assert.Equal(t, errs.Unauthorized, err)
}
