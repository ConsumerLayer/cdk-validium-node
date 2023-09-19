// Code generated by mockery v2.28.1. DO NOT EDIT.

package synchronizer

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	executor "github.com/0xPolygon/cdk-validium-node/state/runtime/executor"

	metrics "github.com/0xPolygon/cdk-validium-node/state/metrics"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/0xPolygon/cdk-validium-node/state"

	types "github.com/ethereum/go-ethereum/core/types"
)

// stateMock is an autogenerated mock type for the stateInterface type
type stateMock struct {
	mock.Mock
}

// AddAccumulatedInputHash provides a mock function with given fields: ctx, batchNum, accInputHash, dbTx
func (_m *stateMock) AddAccumulatedInputHash(ctx context.Context, batchNum uint64, accInputHash common.Hash, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNum, accInputHash, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Hash, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNum, accInputHash, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddBlock provides a mock function with given fields: ctx, block, dbTx
func (_m *stateMock) AddBlock(ctx context.Context, block *state.Block, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, block, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Block, pgx.Tx) error); ok {
		r0 = rf(ctx, block, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddForcedBatch provides a mock function with given fields: ctx, forcedBatch, dbTx
func (_m *stateMock) AddForcedBatch(ctx context.Context, forcedBatch *state.ForcedBatch, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, forcedBatch, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.ForcedBatch, pgx.Tx) error); ok {
		r0 = rf(ctx, forcedBatch, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddForkIDInterval provides a mock function with given fields: ctx, newForkID, dbTx
func (_m *stateMock) AddForkIDInterval(ctx context.Context, newForkID state.ForkIDInterval, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, newForkID, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ForkIDInterval, pgx.Tx) error); ok {
		r0 = rf(ctx, newForkID, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddGlobalExitRoot provides a mock function with given fields: ctx, exitRoot, dbTx
func (_m *stateMock) AddGlobalExitRoot(ctx context.Context, exitRoot *state.GlobalExitRoot, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, exitRoot, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.GlobalExitRoot, pgx.Tx) error); ok {
		r0 = rf(ctx, exitRoot, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddSequence provides a mock function with given fields: ctx, sequence, dbTx
func (_m *stateMock) AddSequence(ctx context.Context, sequence state.Sequence, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, sequence, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.Sequence, pgx.Tx) error); ok {
		r0 = rf(ctx, sequence, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddTrustedReorg provides a mock function with given fields: ctx, trustedReorg, dbTx
func (_m *stateMock) AddTrustedReorg(ctx context.Context, trustedReorg *state.TrustedReorg, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, trustedReorg, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.TrustedReorg, pgx.Tx) error); ok {
		r0 = rf(ctx, trustedReorg, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddVerifiedBatch provides a mock function with given fields: ctx, verifiedBatch, dbTx
func (_m *stateMock) AddVerifiedBatch(ctx context.Context, verifiedBatch *state.VerifiedBatch, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, verifiedBatch, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.VerifiedBatch, pgx.Tx) error); ok {
		r0 = rf(ctx, verifiedBatch, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddVirtualBatch provides a mock function with given fields: ctx, virtualBatch, dbTx
func (_m *stateMock) AddVirtualBatch(ctx context.Context, virtualBatch *state.VirtualBatch, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, virtualBatch, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.VirtualBatch, pgx.Tx) error); ok {
		r0 = rf(ctx, virtualBatch, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *stateMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
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

// CloseBatch provides a mock function with given fields: ctx, receipt, dbTx
func (_m *stateMock) CloseBatch(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, receipt, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingReceipt, pgx.Tx) error); ok {
		r0 = rf(ctx, receipt, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExecuteBatch provides a mock function with given fields: ctx, batch, updateMerkleTree, dbTx
func (_m *stateMock) ExecuteBatch(ctx context.Context, batch state.Batch, updateMerkleTree bool, dbTx pgx.Tx) (*executor.ProcessBatchResponse, error) {
	ret := _m.Called(ctx, batch, updateMerkleTree, dbTx)

	var r0 *executor.ProcessBatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, state.Batch, bool, pgx.Tx) (*executor.ProcessBatchResponse, error)); ok {
		return rf(ctx, batch, updateMerkleTree, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.Batch, bool, pgx.Tx) *executor.ProcessBatchResponse); ok {
		r0 = rf(ctx, batch, updateMerkleTree, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*executor.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.Batch, bool, pgx.Tx) error); ok {
		r1 = rf(ctx, batch, updateMerkleTree, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Batch, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchL2DataByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) GetBatchL2DataByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]byte, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) ([]byte, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []byte); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForkIDByBatchNumber provides a mock function with given fields: batchNumber
func (_m *stateMock) GetForkIDByBatchNumber(batchNumber uint64) uint64 {
	ret := _m.Called(batchNumber)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(batchNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetForkIDs provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetForkIDs(ctx context.Context, dbTx pgx.Tx) ([]state.ForkIDInterval, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 []state.ForkIDInterval
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) ([]state.ForkIDInterval, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) []state.ForkIDInterval); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.ForkIDInterval)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBlock provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVerifiedBatch provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastVerifiedBatch(ctx context.Context, dbTx pgx.Tx) (*state.VerifiedBatch, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.VerifiedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.VerifiedBatch, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.VerifiedBatch); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VerifiedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVirtualBatchNum provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (uint64, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNextForcedBatches provides a mock function with given fields: ctx, nextForcedBatches, dbTx
func (_m *stateMock) GetNextForcedBatches(ctx context.Context, nextForcedBatches int, dbTx pgx.Tx) ([]state.ForcedBatch, error) {
	ret := _m.Called(ctx, nextForcedBatches, dbTx)

	var r0 []state.ForcedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, pgx.Tx) ([]state.ForcedBatch, error)); ok {
		return rf(ctx, nextForcedBatches, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, pgx.Tx) []state.ForcedBatch); ok {
		r0 = rf(ctx, nextForcedBatches, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.ForcedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, pgx.Tx) error); ok {
		r1 = rf(ctx, nextForcedBatches, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPreviousBlock provides a mock function with given fields: ctx, offset, dbTx
func (_m *stateMock) GetPreviousBlock(ctx context.Context, offset uint64, dbTx pgx.Tx) (*state.Block, error) {
	ret := _m.Called(ctx, offset, dbTx)

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Block, error)); ok {
		return rf(ctx, offset, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Block); ok {
		r0 = rf(ctx, offset, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, offset, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReorgedTransactions provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) GetReorgedTransactions(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]*types.Transaction, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 []*types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) ([]*types.Transaction, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []*types.Transaction); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateRootByBatchNumber provides a mock function with given fields: ctx, batchNum, dbTx
func (_m *stateMock) GetStateRootByBatchNumber(ctx context.Context, batchNum uint64, dbTx pgx.Tx) (common.Hash, error) {
	ret := _m.Called(ctx, batchNum, dbTx)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (common.Hash, error)); ok {
		return rf(ctx, batchNum, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) common.Hash); ok {
		r0 = rf(ctx, batchNum, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNum, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStoredFlushID provides a mock function with given fields: ctx
func (_m *stateMock) GetStoredFlushID(ctx context.Context) (uint64, string, error) {
	ret := _m.Called(ctx)

	var r0 uint64
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) string); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// OpenBatch provides a mock function with given fields: ctx, processingContext, dbTx
func (_m *stateMock) OpenBatch(ctx context.Context, processingContext state.ProcessingContext, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, processingContext, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, pgx.Tx) error); ok {
		r0 = rf(ctx, processingContext, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessAndStoreClosedBatch provides a mock function with given fields: ctx, processingCtx, encodedTxs, dbTx, caller
func (_m *stateMock) ProcessAndStoreClosedBatch(ctx context.Context, processingCtx state.ProcessingContext, encodedTxs []byte, dbTx pgx.Tx, caller metrics.CallerLabel) (common.Hash, uint64, string, error) {
	ret := _m.Called(ctx, processingCtx, encodedTxs, dbTx, caller)

	var r0 common.Hash
	var r1 uint64
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, []byte, pgx.Tx, metrics.CallerLabel) (common.Hash, uint64, string, error)); ok {
		return rf(ctx, processingCtx, encodedTxs, dbTx, caller)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, []byte, pgx.Tx, metrics.CallerLabel) common.Hash); ok {
		r0 = rf(ctx, processingCtx, encodedTxs, dbTx, caller)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.ProcessingContext, []byte, pgx.Tx, metrics.CallerLabel) uint64); ok {
		r1 = rf(ctx, processingCtx, encodedTxs, dbTx, caller)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, state.ProcessingContext, []byte, pgx.Tx, metrics.CallerLabel) string); ok {
		r2 = rf(ctx, processingCtx, encodedTxs, dbTx, caller)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(context.Context, state.ProcessingContext, []byte, pgx.Tx, metrics.CallerLabel) error); ok {
		r3 = rf(ctx, processingCtx, encodedTxs, dbTx, caller)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// ProcessBatch provides a mock function with given fields: ctx, request, updateMerkleTree
func (_m *stateMock) ProcessBatch(ctx context.Context, request state.ProcessRequest, updateMerkleTree bool) (*state.ProcessBatchResponse, error) {
	ret := _m.Called(ctx, request, updateMerkleTree)

	var r0 *state.ProcessBatchResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessRequest, bool) (*state.ProcessBatchResponse, error)); ok {
		return rf(ctx, request, updateMerkleTree)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessRequest, bool) *state.ProcessBatchResponse); ok {
		r0 = rf(ctx, request, updateMerkleTree)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProcessBatchResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.ProcessRequest, bool) error); ok {
		r1 = rf(ctx, request, updateMerkleTree)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reset provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *stateMock) Reset(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetForkID provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) ResetForkID(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetTrustedState provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) ResetTrustedState(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetGenesis provides a mock function with given fields: ctx, block, genesis, dbTx
func (_m *stateMock) SetGenesis(ctx context.Context, block state.Block, genesis state.Genesis, dbTx pgx.Tx) ([]byte, error) {
	ret := _m.Called(ctx, block, genesis, dbTx)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, state.Block, state.Genesis, pgx.Tx) ([]byte, error)); ok {
		return rf(ctx, block, genesis, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.Block, state.Genesis, pgx.Tx) []byte); ok {
		r0 = rf(ctx, block, genesis, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.Block, state.Genesis, pgx.Tx) error); ok {
		r1 = rf(ctx, block, genesis, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetInitSyncBatch provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) SetInitSyncBatch(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetLastBatchInfoSeenOnEthereum provides a mock function with given fields: ctx, lastBatchNumberSeen, lastBatchNumberVerified, dbTx
func (_m *stateMock) SetLastBatchInfoSeenOnEthereum(ctx context.Context, lastBatchNumberSeen uint64, lastBatchNumberVerified uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, lastBatchNumberSeen, lastBatchNumberVerified, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, lastBatchNumberSeen, lastBatchNumberVerified, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreTransaction provides a mock function with given fields: ctx, batchNumber, processedTx, coinbase, timestamp, dbTx
func (_m *stateMock) StoreTransaction(ctx context.Context, batchNumber uint64, processedTx *state.ProcessTransactionResponse, coinbase common.Address, timestamp uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, processedTx, coinbase, timestamp, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *state.ProcessTransactionResponse, common.Address, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, processedTx, coinbase, timestamp, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBatchL2Data provides a mock function with given fields: ctx, batchNumber, batchL2Data, dbTx
func (_m *stateMock) UpdateBatchL2Data(ctx context.Context, batchNumber uint64, batchL2Data []byte, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, batchL2Data, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []byte, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, batchL2Data, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewStateMock interface {
	mock.TestingT
	Cleanup(func())
}

// newStateMock creates a new instance of stateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newStateMock(t mockConstructorTestingTnewStateMock) *stateMock {
	mock := &stateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
