// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/ethereum/go-ethereum/core/types"
	mock "github.com/stretchr/testify/mock"
)

// ExternalTxManager is an autogenerated mock type for the ExternalTxManager type
type ExternalTxManager struct {
	mock.Mock
}

// Send provides a mock function with given fields: ctx, tx
func (_m *ExternalTxManager) Send(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	ret := _m.Called(ctx, tx)

	var r0 *types.Receipt
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) (*types.Receipt, error)); ok {
		return rf(ctx, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) *types.Receipt); ok {
		r0 = rf(ctx, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction) error); ok {
		r1 = rf(ctx, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewExternalTxManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewExternalTxManager creates a new instance of ExternalTxManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExternalTxManager(t mockConstructorTestingTNewExternalTxManager) *ExternalTxManager {
	mock := &ExternalTxManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}