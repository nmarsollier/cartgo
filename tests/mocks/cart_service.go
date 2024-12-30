// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/cart/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	cart "github.com/nmarsollier/cartgo/internal/cart"
)

// MockCartService is a mock of CartService interface.
type MockCartService struct {
	ctrl     *gomock.Controller
	recorder *MockCartServiceMockRecorder
}

// MockCartServiceMockRecorder is the mock recorder for MockCartService.
type MockCartServiceMockRecorder struct {
	mock *MockCartService
}

// NewMockCartService creates a new mock instance.
func NewMockCartService(ctrl *gomock.Controller) *MockCartService {
	mock := &MockCartService{ctrl: ctrl}
	mock.recorder = &MockCartServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartService) EXPECT() *MockCartServiceMockRecorder {
	return m.recorder
}

// AddArticle mocks base method.
func (m *MockCartService) AddArticle(userId, articleId string, quantity int) (*cart.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddArticle", userId, articleId, quantity)
	ret0, _ := ret[0].(*cart.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddArticle indicates an expected call of AddArticle.
func (mr *MockCartServiceMockRecorder) AddArticle(userId, articleId, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddArticle", reflect.TypeOf((*MockCartService)(nil).AddArticle), userId, articleId, quantity)
}

// CurrentCart mocks base method.
func (m *MockCartService) CurrentCart(userId string) (*cart.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentCart", userId)
	ret0, _ := ret[0].(*cart.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentCart indicates an expected call of CurrentCart.
func (mr *MockCartServiceMockRecorder) CurrentCart(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentCart", reflect.TypeOf((*MockCartService)(nil).CurrentCart), userId)
}

// FindCartById mocks base method.
func (m *MockCartService) FindCartById(cartId string) (*cart.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCartById", cartId)
	ret0, _ := ret[0].(*cart.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindCartById indicates an expected call of FindCartById.
func (mr *MockCartServiceMockRecorder) FindCartById(cartId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCartById", reflect.TypeOf((*MockCartService)(nil).FindCartById), cartId)
}

// InvalidateCurrentCart mocks base method.
func (m *MockCartService) InvalidateCurrentCart(cry *cart.Cart) (*cart.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateCurrentCart", cry)
	ret0, _ := ret[0].(*cart.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvalidateCurrentCart indicates an expected call of InvalidateCurrentCart.
func (mr *MockCartServiceMockRecorder) InvalidateCurrentCart(cry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateCurrentCart", reflect.TypeOf((*MockCartService)(nil).InvalidateCurrentCart), cry)
}

// ProcessArticleData mocks base method.
func (m *MockCartService) ProcessArticleData(data *cart.ValidationEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessArticleData", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessArticleData indicates an expected call of ProcessArticleData.
func (mr *MockCartServiceMockRecorder) ProcessArticleData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessArticleData", reflect.TypeOf((*MockCartService)(nil).ProcessArticleData), data)
}

// ProcessOrderPlaced mocks base method.
func (m *MockCartService) ProcessOrderPlaced(data *cart.OrderPlacedEvent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessOrderPlaced", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessOrderPlaced indicates an expected call of ProcessOrderPlaced.
func (mr *MockCartServiceMockRecorder) ProcessOrderPlaced(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessOrderPlaced", reflect.TypeOf((*MockCartService)(nil).ProcessOrderPlaced), data)
}

// RemoveArticle mocks base method.
func (m *MockCartService) RemoveArticle(userId, articleId string) (*cart.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveArticle", userId, articleId)
	ret0, _ := ret[0].(*cart.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveArticle indicates an expected call of RemoveArticle.
func (mr *MockCartServiceMockRecorder) RemoveArticle(userId, articleId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveArticle", reflect.TypeOf((*MockCartService)(nil).RemoveArticle), userId, articleId)
}