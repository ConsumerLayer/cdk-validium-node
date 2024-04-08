// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import (
	context "context"

	pgx "github.com/jackc/pgx/v4"
	mock "github.com/stretchr/testify/mock"
)

// StateBeginTransactionInterface is an autogenerated mock type for the StateBeginTransactionInterface type
type StateBeginTransactionInterface struct {
	mock.Mock
}

type StateBeginTransactionInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *StateBeginTransactionInterface) EXPECT() *StateBeginTransactionInterface_Expecter {
	return &StateBeginTransactionInterface_Expecter{mock: &_m.Mock}
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *StateBeginTransactionInterface) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateBeginTransactionInterface_BeginStateTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginStateTransaction'
type StateBeginTransactionInterface_BeginStateTransaction_Call struct {
	*mock.Call
}

// BeginStateTransaction is a helper method to define mock.On call
//   - ctx context.Context
func (_e *StateBeginTransactionInterface_Expecter) BeginStateTransaction(ctx interface{}) *StateBeginTransactionInterface_BeginStateTransaction_Call {
	return &StateBeginTransactionInterface_BeginStateTransaction_Call{Call: _e.mock.On("BeginStateTransaction", ctx)}
}

func (_c *StateBeginTransactionInterface_BeginStateTransaction_Call) Run(run func(ctx context.Context)) *StateBeginTransactionInterface_BeginStateTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *StateBeginTransactionInterface_BeginStateTransaction_Call) Return(_a0 pgx.Tx, _a1 error) *StateBeginTransactionInterface_BeginStateTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateBeginTransactionInterface_BeginStateTransaction_Call) RunAndReturn(run func(context.Context) (pgx.Tx, error)) *StateBeginTransactionInterface_BeginStateTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewStateBeginTransactionInterface creates a new instance of StateBeginTransactionInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateBeginTransactionInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateBeginTransactionInterface {
	mock := &StateBeginTransactionInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
