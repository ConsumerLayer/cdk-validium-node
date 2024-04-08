// Code generated by mockery v2.32.0. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/0xPolygonHermez/zkevm-node/state"
)

// StateMock is an autogenerated mock type for the stateInterface type
type StateMock struct {
	mock.Mock
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *StateMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
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

// BuildChangeL2Block provides a mock function with given fields: deltaTimestamp, l1InfoTreeIndex
func (_m *StateMock) BuildChangeL2Block(deltaTimestamp uint32, l1InfoTreeIndex uint32) []byte {
	ret := _m.Called(deltaTimestamp, l1InfoTreeIndex)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(uint32, uint32) []byte); ok {
		r0 = rf(deltaTimestamp, l1InfoTreeIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// CloseBatch provides a mock function with given fields: ctx, receipt, dbTx
func (_m *StateMock) CloseBatch(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, receipt, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingReceipt, pgx.Tx) error); ok {
		r0 = rf(ctx, receipt, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CloseWIPBatch provides a mock function with given fields: ctx, receipt, dbTx
func (_m *StateMock) CloseWIPBatch(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, receipt, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingReceipt, pgx.Tx) error); ok {
		r0 = rf(ctx, receipt, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountReorgs provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) CountReorgs(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
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

// GetBalanceByStateRoot provides a mock function with given fields: ctx, address, root
func (_m *StateMock) GetBalanceByStateRoot(ctx context.Context, address common.Address, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) (*big.Int, error)); ok {
		return rf(ctx, address, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
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

// GetBlockByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *StateMock) GetBlockByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (*state.Block, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 *state.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.Block, error)); ok {
		return rf(ctx, blockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Block); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDSBatches provides a mock function with given fields: ctx, firstBatchNumber, lastBatchNumber, readWIPBatch, dbTx
func (_m *StateMock) GetDSBatches(ctx context.Context, firstBatchNumber uint64, lastBatchNumber uint64, readWIPBatch bool, dbTx pgx.Tx) ([]*state.DSBatch, error) {
	ret := _m.Called(ctx, firstBatchNumber, lastBatchNumber, readWIPBatch, dbTx)

	var r0 []*state.DSBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, bool, pgx.Tx) ([]*state.DSBatch, error)); ok {
		return rf(ctx, firstBatchNumber, lastBatchNumber, readWIPBatch, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, bool, pgx.Tx) []*state.DSBatch); ok {
		r0 = rf(ctx, firstBatchNumber, lastBatchNumber, readWIPBatch, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.DSBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, bool, pgx.Tx) error); ok {
		r1 = rf(ctx, firstBatchNumber, lastBatchNumber, readWIPBatch, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDSGenesisBlock provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetDSGenesisBlock(ctx context.Context, dbTx pgx.Tx) (*state.DSL2Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.DSL2Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.DSL2Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.DSL2Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.DSL2Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDSL2Blocks provides a mock function with given fields: ctx, firstBatchNumber, lastBatchNumber, dbTx
func (_m *StateMock) GetDSL2Blocks(ctx context.Context, firstBatchNumber uint64, lastBatchNumber uint64, dbTx pgx.Tx) ([]*state.DSL2Block, error) {
	ret := _m.Called(ctx, firstBatchNumber, lastBatchNumber, dbTx)

	var r0 []*state.DSL2Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) ([]*state.DSL2Block, error)); ok {
		return rf(ctx, firstBatchNumber, lastBatchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) []*state.DSL2Block); ok {
		r0 = rf(ctx, firstBatchNumber, lastBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.DSL2Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, firstBatchNumber, lastBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDSL2Transactions provides a mock function with given fields: ctx, firstL2Block, lastL2Block, dbTx
func (_m *StateMock) GetDSL2Transactions(ctx context.Context, firstL2Block uint64, lastL2Block uint64, dbTx pgx.Tx) ([]*state.DSL2Transaction, error) {
	ret := _m.Called(ctx, firstL2Block, lastL2Block, dbTx)

	var r0 []*state.DSL2Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) ([]*state.DSL2Transaction, error)); ok {
		return rf(ctx, firstL2Block, lastL2Block, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) []*state.DSL2Transaction); ok {
		r0 = rf(ctx, firstL2Block, lastL2Block, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.DSL2Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, firstL2Block, lastL2Block, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForcedBatch provides a mock function with given fields: ctx, forcedBatchNumber, dbTx
func (_m *StateMock) GetForcedBatch(ctx context.Context, forcedBatchNumber uint64, dbTx pgx.Tx) (*state.ForcedBatch, error) {
	ret := _m.Called(ctx, forcedBatchNumber, dbTx)

	var r0 *state.ForcedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (*state.ForcedBatch, error)); ok {
		return rf(ctx, forcedBatchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.ForcedBatch); ok {
		r0 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ForcedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForcedBatchParentHash provides a mock function with given fields: ctx, forcedBatchNumber, dbTx
func (_m *StateMock) GetForcedBatchParentHash(ctx context.Context, forcedBatchNumber uint64, dbTx pgx.Tx) (common.Hash, error) {
	ret := _m.Called(ctx, forcedBatchNumber, dbTx)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (common.Hash, error)); ok {
		return rf(ctx, forcedBatchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) common.Hash); ok {
		r0 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForcedBatchesSince provides a mock function with given fields: ctx, forcedBatchNumber, maxBlockNumber, dbTx
func (_m *StateMock) GetForcedBatchesSince(ctx context.Context, forcedBatchNumber uint64, maxBlockNumber uint64, dbTx pgx.Tx) ([]*state.ForcedBatch, error) {
	ret := _m.Called(ctx, forcedBatchNumber, maxBlockNumber, dbTx)

	var r0 []*state.ForcedBatch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) ([]*state.ForcedBatch, error)); ok {
		return rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) []*state.ForcedBatch); ok {
		r0 = rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.ForcedBatch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, forcedBatchNumber, maxBlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForkIDByBatchNumber provides a mock function with given fields: batchNumber
func (_m *StateMock) GetForkIDByBatchNumber(batchNumber uint64) uint64 {
	ret := _m.Called(batchNumber)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(batchNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetL1InfoRootLeafByIndex provides a mock function with given fields: ctx, l1InfoTreeIndex, dbTx
func (_m *StateMock) GetL1InfoRootLeafByIndex(ctx context.Context, l1InfoTreeIndex uint32, dbTx pgx.Tx) (state.L1InfoTreeExitRootStorageEntry, error) {
	ret := _m.Called(ctx, l1InfoTreeIndex, dbTx)

	var r0 state.L1InfoTreeExitRootStorageEntry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, pgx.Tx) (state.L1InfoTreeExitRootStorageEntry, error)); ok {
		return rf(ctx, l1InfoTreeIndex, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, pgx.Tx) state.L1InfoTreeExitRootStorageEntry); ok {
		r0 = rf(ctx, l1InfoTreeIndex, dbTx)
	} else {
		r0 = ret.Get(0).(state.L1InfoTreeExitRootStorageEntry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, pgx.Tx) error); ok {
		r1 = rf(ctx, l1InfoTreeIndex, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL1InfoTreeDataFromBatchL2Data provides a mock function with given fields: ctx, batchL2Data, dbTx
func (_m *StateMock) GetL1InfoTreeDataFromBatchL2Data(ctx context.Context, batchL2Data []byte, dbTx pgx.Tx) (map[uint32]state.L1DataV2, common.Hash, common.Hash, error) {
	ret := _m.Called(ctx, batchL2Data, dbTx)

	var r0 map[uint32]state.L1DataV2
	var r1 common.Hash
	var r2 common.Hash
	var r3 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, pgx.Tx) (map[uint32]state.L1DataV2, common.Hash, common.Hash, error)); ok {
		return rf(ctx, batchL2Data, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, pgx.Tx) map[uint32]state.L1DataV2); ok {
		r0 = rf(ctx, batchL2Data, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]state.L1DataV2)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, pgx.Tx) common.Hash); ok {
		r1 = rf(ctx, batchL2Data, dbTx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(common.Hash)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, []byte, pgx.Tx) common.Hash); ok {
		r2 = rf(ctx, batchL2Data, dbTx)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(common.Hash)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, []byte, pgx.Tx) error); ok {
		r3 = rf(ctx, batchL2Data, dbTx)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetLastBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
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
func (_m *StateMock) GetLastBlock(ctx context.Context, dbTx pgx.Tx) (*state.Block, error) {
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

// GetLastL2Block provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastL2Block(ctx context.Context, dbTx pgx.Tx) (*state.L2Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.L2Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (*state.L2Block, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.L2Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.L2Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastStateRoot provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastStateRoot(ctx context.Context, dbTx pgx.Tx) (common.Hash, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (common.Hash, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) common.Hash); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastTrustedForcedBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastTrustedForcedBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
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

// GetLastVirtualBatchNum provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
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

// GetLatestBatchGlobalExitRoot provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLatestBatchGlobalExitRoot(ctx context.Context, dbTx pgx.Tx) (common.Hash, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) (common.Hash, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) common.Hash); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestL1InfoRoot provides a mock function with given fields: ctx, maxBlockNumber
func (_m *StateMock) GetLatestL1InfoRoot(ctx context.Context, maxBlockNumber uint64) (state.L1InfoTreeExitRootStorageEntry, error) {
	ret := _m.Called(ctx, maxBlockNumber)

	var r0 state.L1InfoTreeExitRootStorageEntry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (state.L1InfoTreeExitRootStorageEntry, error)); ok {
		return rf(ctx, maxBlockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) state.L1InfoTreeExitRootStorageEntry); ok {
		r0 = rf(ctx, maxBlockNumber)
	} else {
		r0 = ret.Get(0).(state.L1InfoTreeExitRootStorageEntry)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, maxBlockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonceByStateRoot provides a mock function with given fields: ctx, address, root
func (_m *StateMock) GetNonceByStateRoot(ctx context.Context, address common.Address, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) (*big.Int, error)); ok {
		return rf(ctx, address, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, common.Hash) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNotCheckedBatches provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetNotCheckedBatches(ctx context.Context, dbTx pgx.Tx) ([]*state.Batch, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 []*state.Batch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) ([]*state.Batch, error)); ok {
		return rf(ctx, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) []*state.Batch); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*state.Batch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageAt provides a mock function with given fields: ctx, address, position, root
func (_m *StateMock) GetStorageAt(ctx context.Context, address common.Address, position *big.Int, root common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, address, position, root)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, common.Hash) (*big.Int, error)); ok {
		return rf(ctx, address, position, root)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, common.Hash) *big.Int); ok {
		r0 = rf(ctx, address, position, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int, common.Hash) error); ok {
		r1 = rf(ctx, address, position, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStoredFlushID provides a mock function with given fields: ctx
func (_m *StateMock) GetStoredFlushID(ctx context.Context) (uint64, string, error) {
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

// GetTxsOlderThanNL1BlocksUntilTxHash provides a mock function with given fields: ctx, nL1Blocks, earliestTxHash, dbTx
func (_m *StateMock) GetTxsOlderThanNL1BlocksUntilTxHash(ctx context.Context, nL1Blocks uint64, earliestTxHash common.Hash, dbTx pgx.Tx) ([]common.Hash, error) {
	ret := _m.Called(ctx, nL1Blocks, earliestTxHash, dbTx)

	var r0 []common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Hash, pgx.Tx) ([]common.Hash, error)); ok {
		return rf(ctx, nL1Blocks, earliestTxHash, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, common.Hash, pgx.Tx) []common.Hash); ok {
		r0 = rf(ctx, nL1Blocks, earliestTxHash, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, nL1Blocks, earliestTxHash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVirtualBatchParentHash provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetVirtualBatchParentHash(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (common.Hash, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) (common.Hash, error)); ok {
		return rf(ctx, batchNumber, dbTx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) common.Hash); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OpenBatch provides a mock function with given fields: ctx, processingContext, dbTx
func (_m *StateMock) OpenBatch(ctx context.Context, processingContext state.ProcessingContext, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, processingContext, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingContext, pgx.Tx) error); ok {
		r0 = rf(ctx, processingContext, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OpenWIPBatch provides a mock function with given fields: ctx, batch, dbTx
func (_m *StateMock) OpenWIPBatch(ctx context.Context, batch state.Batch, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batch, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.Batch, pgx.Tx) error); ok {
		r0 = rf(ctx, batch, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessBatchV2 provides a mock function with given fields: ctx, request, updateMerkleTree
func (_m *StateMock) ProcessBatchV2(ctx context.Context, request state.ProcessRequest, updateMerkleTree bool) (*state.ProcessBatchResponse, error) {
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

// StoreL2Block provides a mock function with given fields: ctx, batchNumber, l2Block, txsEGPLog, dbTx
func (_m *StateMock) StoreL2Block(ctx context.Context, batchNumber uint64, l2Block *state.ProcessBlockResponse, txsEGPLog []*state.EffectiveGasPriceLog, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, l2Block, txsEGPLog, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *state.ProcessBlockResponse, []*state.EffectiveGasPriceLog, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, l2Block, txsEGPLog, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBatchAsChecked provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) UpdateBatchAsChecked(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) error); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWIPBatch provides a mock function with given fields: ctx, receipt, dbTx
func (_m *StateMock) UpdateWIPBatch(ctx context.Context, receipt state.ProcessingReceipt, dbTx pgx.Tx) error {
	ret := _m.Called(ctx, receipt, dbTx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessingReceipt, pgx.Tx) error); ok {
		r0 = rf(ctx, receipt, dbTx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStateMock creates a new instance of StateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateMock {
	mock := &StateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
