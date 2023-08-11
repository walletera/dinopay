// Code generated by mockery v2.32.4. DO NOT EDIT.

package tests

import (
	context "context"

	api "github.com/walletera/dinopay/api"

	mock "github.com/stretchr/testify/mock"
)

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

type MockHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHandler) EXPECT() *MockHandler_Expecter {
	return &MockHandler_Expecter{mock: &_m.Mock}
}

// CreateEventSubscription provides a mock function with given fields: ctx, req
func (_m *MockHandler) CreateEventSubscription(ctx context.Context, req *api.EventSubscription) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.EventSubscription) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockHandler_CreateEventSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEventSubscription'
type MockHandler_CreateEventSubscription_Call struct {
	*mock.Call
}

// CreateEventSubscription is a helper method to define mock.On call
//   - ctx context.Context
//   - req *api.EventSubscription
func (_e *MockHandler_Expecter) CreateEventSubscription(ctx interface{}, req interface{}) *MockHandler_CreateEventSubscription_Call {
	return &MockHandler_CreateEventSubscription_Call{Call: _e.mock.On("CreateEventSubscription", ctx, req)}
}

func (_c *MockHandler_CreateEventSubscription_Call) Run(run func(ctx context.Context, req *api.EventSubscription)) *MockHandler_CreateEventSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*api.EventSubscription))
	})
	return _c
}

func (_c *MockHandler_CreateEventSubscription_Call) Return(_a0 error) *MockHandler_CreateEventSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHandler_CreateEventSubscription_Call) RunAndReturn(run func(context.Context, *api.EventSubscription) error) *MockHandler_CreateEventSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// CreatePayment provides a mock function with given fields: ctx, req
func (_m *MockHandler) CreatePayment(ctx context.Context, req *api.Payment) (api.CreatePaymentRes, error) {
	ret := _m.Called(ctx, req)

	var r0 api.CreatePaymentRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *api.Payment) (api.CreatePaymentRes, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *api.Payment) api.CreatePaymentRes); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.CreatePaymentRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *api.Payment) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHandler_CreatePayment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreatePayment'
type MockHandler_CreatePayment_Call struct {
	*mock.Call
}

// CreatePayment is a helper method to define mock.On call
//   - ctx context.Context
//   - req *api.Payment
func (_e *MockHandler_Expecter) CreatePayment(ctx interface{}, req interface{}) *MockHandler_CreatePayment_Call {
	return &MockHandler_CreatePayment_Call{Call: _e.mock.On("CreatePayment", ctx, req)}
}

func (_c *MockHandler_CreatePayment_Call) Run(run func(ctx context.Context, req *api.Payment)) *MockHandler_CreatePayment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*api.Payment))
	})
	return _c
}

func (_c *MockHandler_CreatePayment_Call) Return(_a0 api.CreatePaymentRes, _a1 error) *MockHandler_CreatePayment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHandler_CreatePayment_Call) RunAndReturn(run func(context.Context, *api.Payment) (api.CreatePaymentRes, error)) *MockHandler_CreatePayment_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHandler creates a new instance of MockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandler {
	mock := &MockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}