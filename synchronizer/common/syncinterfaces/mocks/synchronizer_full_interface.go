// Code generated by mockery. DO NOT EDIT.

package mock_syncinterfaces

import (
	pgx "github.com/jackc/pgx/v4"
	mock "github.com/stretchr/testify/mock"
)

// SynchronizerFullInterface is an autogenerated mock type for the SynchronizerFullInterface type
type SynchronizerFullInterface struct {
	mock.Mock
}

type SynchronizerFullInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *SynchronizerFullInterface) EXPECT() *SynchronizerFullInterface_Expecter {
	return &SynchronizerFullInterface_Expecter{mock: &_m.Mock}
}

// CheckFlushID provides a mock function with given fields: dbTx
func (_m *SynchronizerFullInterface) CheckFlushID(dbTx pgx.Tx) error {
	ret := _m.Called(dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(pgx.Tx) error); ok {
		r0 = rf(dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SynchronizerFullInterface_CheckFlushID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckFlushID'
type SynchronizerFullInterface_CheckFlushID_Call struct {
	*mock.Call
}

// CheckFlushID is a helper method to define mock.On call
//   - dbTx pgx.Tx
func (_e *SynchronizerFullInterface_Expecter) CheckFlushID(dbTx interface{}) *SynchronizerFullInterface_CheckFlushID_Call {
	return &SynchronizerFullInterface_CheckFlushID_Call{Call: _e.mock.On("CheckFlushID", dbTx)}
}

func (_c *SynchronizerFullInterface_CheckFlushID_Call) Run(run func(dbTx pgx.Tx)) *SynchronizerFullInterface_CheckFlushID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(pgx.Tx))
	})
	return _c
}

func (_c *SynchronizerFullInterface_CheckFlushID_Call) Return(_a0 error) *SynchronizerFullInterface_CheckFlushID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SynchronizerFullInterface_CheckFlushID_Call) RunAndReturn(run func(pgx.Tx) error) *SynchronizerFullInterface_CheckFlushID_Call {
	_c.Call.Return(run)
	return _c
}

// CleanTrustedState provides a mock function with given fields:
func (_m *SynchronizerFullInterface) CleanTrustedState() {
	_m.Called()
}

// SynchronizerFullInterface_CleanTrustedState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CleanTrustedState'
type SynchronizerFullInterface_CleanTrustedState_Call struct {
	*mock.Call
}

// CleanTrustedState is a helper method to define mock.On call
func (_e *SynchronizerFullInterface_Expecter) CleanTrustedState() *SynchronizerFullInterface_CleanTrustedState_Call {
	return &SynchronizerFullInterface_CleanTrustedState_Call{Call: _e.mock.On("CleanTrustedState")}
}

func (_c *SynchronizerFullInterface_CleanTrustedState_Call) Run(run func()) *SynchronizerFullInterface_CleanTrustedState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SynchronizerFullInterface_CleanTrustedState_Call) Return() *SynchronizerFullInterface_CleanTrustedState_Call {
	_c.Call.Return()
	return _c
}

func (_c *SynchronizerFullInterface_CleanTrustedState_Call) RunAndReturn(run func()) *SynchronizerFullInterface_CleanTrustedState_Call {
	_c.Call.Return(run)
	return _c
}

// IsTrustedSequencer provides a mock function with given fields:
func (_m *SynchronizerFullInterface) IsTrustedSequencer() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SynchronizerFullInterface_IsTrustedSequencer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsTrustedSequencer'
type SynchronizerFullInterface_IsTrustedSequencer_Call struct {
	*mock.Call
}

// IsTrustedSequencer is a helper method to define mock.On call
func (_e *SynchronizerFullInterface_Expecter) IsTrustedSequencer() *SynchronizerFullInterface_IsTrustedSequencer_Call {
	return &SynchronizerFullInterface_IsTrustedSequencer_Call{Call: _e.mock.On("IsTrustedSequencer")}
}

func (_c *SynchronizerFullInterface_IsTrustedSequencer_Call) Run(run func()) *SynchronizerFullInterface_IsTrustedSequencer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *SynchronizerFullInterface_IsTrustedSequencer_Call) Return(_a0 bool) *SynchronizerFullInterface_IsTrustedSequencer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SynchronizerFullInterface_IsTrustedSequencer_Call) RunAndReturn(run func() bool) *SynchronizerFullInterface_IsTrustedSequencer_Call {
	_c.Call.Return(run)
	return _c
}

// PendingFlushID provides a mock function with given fields: flushID, proverID
func (_m *SynchronizerFullInterface) PendingFlushID(flushID uint64, proverID string) {
	_m.Called(flushID, proverID)
}

// SynchronizerFullInterface_PendingFlushID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingFlushID'
type SynchronizerFullInterface_PendingFlushID_Call struct {
	*mock.Call
}

// PendingFlushID is a helper method to define mock.On call
//   - flushID uint64
//   - proverID string
func (_e *SynchronizerFullInterface_Expecter) PendingFlushID(flushID interface{}, proverID interface{}) *SynchronizerFullInterface_PendingFlushID_Call {
	return &SynchronizerFullInterface_PendingFlushID_Call{Call: _e.mock.On("PendingFlushID", flushID, proverID)}
}

func (_c *SynchronizerFullInterface_PendingFlushID_Call) Run(run func(flushID uint64, proverID string)) *SynchronizerFullInterface_PendingFlushID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint64), args[1].(string))
	})
	return _c
}

func (_c *SynchronizerFullInterface_PendingFlushID_Call) Return() *SynchronizerFullInterface_PendingFlushID_Call {
	_c.Call.Return()
	return _c
}

func (_c *SynchronizerFullInterface_PendingFlushID_Call) RunAndReturn(run func(uint64, string)) *SynchronizerFullInterface_PendingFlushID_Call {
	_c.Call.Return(run)
	return _c
}

// NewSynchronizerFullInterface creates a new instance of SynchronizerFullInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSynchronizerFullInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SynchronizerFullInterface {
	mock := &SynchronizerFullInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
