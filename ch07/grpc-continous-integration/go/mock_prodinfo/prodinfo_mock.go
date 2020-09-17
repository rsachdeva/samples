// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/grpc-up-and-running/samples/ch07/grpc-docker/go/proto-gen (interfaces: ProductInfoClient)

// Package mock_proto_gen is a generated GoMock package.
package mock_proto_gen

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	ecommerce "github.com/grpc-up-and-running/samples/ch07/grpc-docker/go/proto-gen"
	grpc "google.golang.org/grpc"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

// MockProductInfoClient is a mock of ProductInfoClient interface
type MockProductInfoClient struct {
	ctrl     *gomock.Controller
	recorder *MockProductInfoClientMockRecorder
}

// MockProductInfoClientMockRecorder is the mock recorder for MockProductInfoClient
type MockProductInfoClientMockRecorder struct {
	mock *MockProductInfoClient
}

// NewMockProductInfoClient creates a new mock instance
func NewMockProductInfoClient(ctrl *gomock.Controller) *MockProductInfoClient {
	mock := &MockProductInfoClient{ctrl: ctrl}
	mock.recorder = &MockProductInfoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProductInfoClient) EXPECT() *MockProductInfoClientMockRecorder {
	return m.recorder
}

// AddProduct mocks base method
func (m *MockProductInfoClient) AddProduct(arg0 context.Context, arg1 *ecommerce.Product, arg2 ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProduct", varargs...)
	ret0, _ := ret[0].(*wrapperspb.StringValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct
func (mr *MockProductInfoClientMockRecorder) AddProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockProductInfoClient)(nil).AddProduct), varargs...)
}

// GetProduct mocks base method
func (m *MockProductInfoClient) GetProduct(arg0 context.Context, arg1 *wrapperspb.StringValue, arg2 ...grpc.CallOption) (*ecommerce.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProduct", varargs...)
	ret0, _ := ret[0].(*ecommerce.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct
func (mr *MockProductInfoClientMockRecorder) GetProduct(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductInfoClient)(nil).GetProduct), varargs...)
}
