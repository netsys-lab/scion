// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/lib/pathdb (interfaces: PathDB,Transaction,ReadWrite)

// Package mock_pathdb is a generated GoMock package.
package mock_pathdb

import (
	context "context"
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	addr "github.com/scionproto/scion/go/lib/addr"
	seg "github.com/scionproto/scion/go/lib/ctrl/seg"
	pathdb "github.com/scionproto/scion/go/lib/pathdb"
	query "github.com/scionproto/scion/go/lib/pathdb/query"
	reflect "reflect"
	time "time"
)

// MockPathDB is a mock of PathDB interface
type MockPathDB struct {
	ctrl     *gomock.Controller
	recorder *MockPathDBMockRecorder
}

// MockPathDBMockRecorder is the mock recorder for MockPathDB
type MockPathDBMockRecorder struct {
	mock *MockPathDB
}

// NewMockPathDB creates a new mock instance
func NewMockPathDB(ctrl *gomock.Controller) *MockPathDB {
	mock := &MockPathDB{ctrl: ctrl}
	mock.recorder = &MockPathDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPathDB) EXPECT() *MockPathDBMockRecorder {
	return m.recorder
}

// BeginTransaction mocks base method
func (m *MockPathDB) BeginTransaction(arg0 context.Context, arg1 *sql.TxOptions) (pathdb.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTransaction", arg0, arg1)
	ret0, _ := ret[0].(pathdb.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTransaction indicates an expected call of BeginTransaction
func (mr *MockPathDBMockRecorder) BeginTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTransaction", reflect.TypeOf((*MockPathDB)(nil).BeginTransaction), arg0, arg1)
}

// Close mocks base method
func (m *MockPathDB) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPathDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPathDB)(nil).Close))
}

// Delete mocks base method
func (m *MockPathDB) Delete(arg0 context.Context, arg1 *query.Params) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockPathDBMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPathDB)(nil).Delete), arg0, arg1)
}

// DeleteExpired mocks base method
func (m *MockPathDB) DeleteExpired(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpired", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpired indicates an expected call of DeleteExpired
func (mr *MockPathDBMockRecorder) DeleteExpired(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpired", reflect.TypeOf((*MockPathDB)(nil).DeleteExpired), arg0, arg1)
}

// DeleteExpiredNQ mocks base method
func (m *MockPathDB) DeleteExpiredNQ(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredNQ", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpiredNQ indicates an expected call of DeleteExpiredNQ
func (mr *MockPathDBMockRecorder) DeleteExpiredNQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredNQ", reflect.TypeOf((*MockPathDB)(nil).DeleteExpiredNQ), arg0, arg1)
}

// DeleteNQ mocks base method
func (m *MockPathDB) DeleteNQ(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNQ", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNQ indicates an expected call of DeleteNQ
func (mr *MockPathDBMockRecorder) DeleteNQ(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNQ", reflect.TypeOf((*MockPathDB)(nil).DeleteNQ), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockPathDB) Get(arg0 context.Context, arg1 *query.Params) (query.Results, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(query.Results)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPathDBMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPathDB)(nil).Get), arg0, arg1)
}

// GetAll mocks base method
func (m *MockPathDB) GetAll(arg0 context.Context) (<-chan query.ResultOrErr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(<-chan query.ResultOrErr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockPathDBMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPathDB)(nil).GetAll), arg0)
}

// GetNextQuery mocks base method
func (m *MockPathDB) GetNextQuery(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextQuery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextQuery indicates an expected call of GetNextQuery
func (mr *MockPathDBMockRecorder) GetNextQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextQuery", reflect.TypeOf((*MockPathDB)(nil).GetNextQuery), arg0, arg1, arg2, arg3)
}

// Insert mocks base method
func (m *MockPathDB) Insert(arg0 context.Context, arg1 *seg.Meta) (pathdb.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(pathdb.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockPathDBMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPathDB)(nil).Insert), arg0, arg1)
}

// InsertNextQuery mocks base method
func (m *MockPathDB) InsertNextQuery(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash, arg4 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNextQuery", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNextQuery indicates an expected call of InsertNextQuery
func (mr *MockPathDBMockRecorder) InsertNextQuery(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNextQuery", reflect.TypeOf((*MockPathDB)(nil).InsertNextQuery), arg0, arg1, arg2, arg3, arg4)
}

// InsertWithHPCfgIDs mocks base method
func (m *MockPathDB) InsertWithHPCfgIDs(arg0 context.Context, arg1 *seg.Meta, arg2 []*query.HPCfgID) (pathdb.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWithHPCfgIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].(pathdb.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertWithHPCfgIDs indicates an expected call of InsertWithHPCfgIDs
func (mr *MockPathDBMockRecorder) InsertWithHPCfgIDs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithHPCfgIDs", reflect.TypeOf((*MockPathDB)(nil).InsertWithHPCfgIDs), arg0, arg1, arg2)
}

// SetMaxIdleConns mocks base method
func (m *MockPathDB) SetMaxIdleConns(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxIdleConns", arg0)
}

// SetMaxIdleConns indicates an expected call of SetMaxIdleConns
func (mr *MockPathDBMockRecorder) SetMaxIdleConns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxIdleConns", reflect.TypeOf((*MockPathDB)(nil).SetMaxIdleConns), arg0)
}

// SetMaxOpenConns mocks base method
func (m *MockPathDB) SetMaxOpenConns(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxOpenConns", arg0)
}

// SetMaxOpenConns indicates an expected call of SetMaxOpenConns
func (mr *MockPathDBMockRecorder) SetMaxOpenConns(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxOpenConns", reflect.TypeOf((*MockPathDB)(nil).SetMaxOpenConns), arg0)
}

// MockTransaction is a mock of Transaction interface
type MockTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionMockRecorder
}

// MockTransactionMockRecorder is the mock recorder for MockTransaction
type MockTransactionMockRecorder struct {
	mock *MockTransaction
}

// NewMockTransaction creates a new mock instance
func NewMockTransaction(ctrl *gomock.Controller) *MockTransaction {
	mock := &MockTransaction{ctrl: ctrl}
	mock.recorder = &MockTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransaction) EXPECT() *MockTransactionMockRecorder {
	return m.recorder
}

// Commit mocks base method
func (m *MockTransaction) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockTransactionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTransaction)(nil).Commit))
}

// Delete mocks base method
func (m *MockTransaction) Delete(arg0 context.Context, arg1 *query.Params) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockTransactionMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTransaction)(nil).Delete), arg0, arg1)
}

// DeleteExpired mocks base method
func (m *MockTransaction) DeleteExpired(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpired", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpired indicates an expected call of DeleteExpired
func (mr *MockTransactionMockRecorder) DeleteExpired(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpired", reflect.TypeOf((*MockTransaction)(nil).DeleteExpired), arg0, arg1)
}

// DeleteExpiredNQ mocks base method
func (m *MockTransaction) DeleteExpiredNQ(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredNQ", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpiredNQ indicates an expected call of DeleteExpiredNQ
func (mr *MockTransactionMockRecorder) DeleteExpiredNQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredNQ", reflect.TypeOf((*MockTransaction)(nil).DeleteExpiredNQ), arg0, arg1)
}

// DeleteNQ mocks base method
func (m *MockTransaction) DeleteNQ(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNQ", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNQ indicates an expected call of DeleteNQ
func (mr *MockTransactionMockRecorder) DeleteNQ(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNQ", reflect.TypeOf((*MockTransaction)(nil).DeleteNQ), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockTransaction) Get(arg0 context.Context, arg1 *query.Params) (query.Results, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(query.Results)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTransactionMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTransaction)(nil).Get), arg0, arg1)
}

// GetAll mocks base method
func (m *MockTransaction) GetAll(arg0 context.Context) (<-chan query.ResultOrErr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(<-chan query.ResultOrErr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockTransactionMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTransaction)(nil).GetAll), arg0)
}

// GetNextQuery mocks base method
func (m *MockTransaction) GetNextQuery(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextQuery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextQuery indicates an expected call of GetNextQuery
func (mr *MockTransactionMockRecorder) GetNextQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextQuery", reflect.TypeOf((*MockTransaction)(nil).GetNextQuery), arg0, arg1, arg2, arg3)
}

// Insert mocks base method
func (m *MockTransaction) Insert(arg0 context.Context, arg1 *seg.Meta) (pathdb.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(pathdb.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockTransactionMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTransaction)(nil).Insert), arg0, arg1)
}

// InsertNextQuery mocks base method
func (m *MockTransaction) InsertNextQuery(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash, arg4 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNextQuery", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNextQuery indicates an expected call of InsertNextQuery
func (mr *MockTransactionMockRecorder) InsertNextQuery(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNextQuery", reflect.TypeOf((*MockTransaction)(nil).InsertNextQuery), arg0, arg1, arg2, arg3, arg4)
}

// InsertWithHPCfgIDs mocks base method
func (m *MockTransaction) InsertWithHPCfgIDs(arg0 context.Context, arg1 *seg.Meta, arg2 []*query.HPCfgID) (pathdb.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWithHPCfgIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].(pathdb.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertWithHPCfgIDs indicates an expected call of InsertWithHPCfgIDs
func (mr *MockTransactionMockRecorder) InsertWithHPCfgIDs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithHPCfgIDs", reflect.TypeOf((*MockTransaction)(nil).InsertWithHPCfgIDs), arg0, arg1, arg2)
}

// Rollback mocks base method
func (m *MockTransaction) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback
func (mr *MockTransactionMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockTransaction)(nil).Rollback))
}

// MockReadWrite is a mock of ReadWrite interface
type MockReadWrite struct {
	ctrl     *gomock.Controller
	recorder *MockReadWriteMockRecorder
}

// MockReadWriteMockRecorder is the mock recorder for MockReadWrite
type MockReadWriteMockRecorder struct {
	mock *MockReadWrite
}

// NewMockReadWrite creates a new mock instance
func NewMockReadWrite(ctrl *gomock.Controller) *MockReadWrite {
	mock := &MockReadWrite{ctrl: ctrl}
	mock.recorder = &MockReadWriteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadWrite) EXPECT() *MockReadWriteMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockReadWrite) Delete(arg0 context.Context, arg1 *query.Params) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockReadWriteMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReadWrite)(nil).Delete), arg0, arg1)
}

// DeleteExpired mocks base method
func (m *MockReadWrite) DeleteExpired(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpired", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpired indicates an expected call of DeleteExpired
func (mr *MockReadWriteMockRecorder) DeleteExpired(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpired", reflect.TypeOf((*MockReadWrite)(nil).DeleteExpired), arg0, arg1)
}

// DeleteExpiredNQ mocks base method
func (m *MockReadWrite) DeleteExpiredNQ(arg0 context.Context, arg1 time.Time) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpiredNQ", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExpiredNQ indicates an expected call of DeleteExpiredNQ
func (mr *MockReadWriteMockRecorder) DeleteExpiredNQ(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpiredNQ", reflect.TypeOf((*MockReadWrite)(nil).DeleteExpiredNQ), arg0, arg1)
}

// DeleteNQ mocks base method
func (m *MockReadWrite) DeleteNQ(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNQ", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNQ indicates an expected call of DeleteNQ
func (mr *MockReadWriteMockRecorder) DeleteNQ(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNQ", reflect.TypeOf((*MockReadWrite)(nil).DeleteNQ), arg0, arg1, arg2, arg3)
}

// Get mocks base method
func (m *MockReadWrite) Get(arg0 context.Context, arg1 *query.Params) (query.Results, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(query.Results)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockReadWriteMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReadWrite)(nil).Get), arg0, arg1)
}

// GetAll mocks base method
func (m *MockReadWrite) GetAll(arg0 context.Context) (<-chan query.ResultOrErr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(<-chan query.ResultOrErr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockReadWriteMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockReadWrite)(nil).GetAll), arg0)
}

// GetNextQuery mocks base method
func (m *MockReadWrite) GetNextQuery(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextQuery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextQuery indicates an expected call of GetNextQuery
func (mr *MockReadWriteMockRecorder) GetNextQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextQuery", reflect.TypeOf((*MockReadWrite)(nil).GetNextQuery), arg0, arg1, arg2, arg3)
}

// Insert mocks base method
func (m *MockReadWrite) Insert(arg0 context.Context, arg1 *seg.Meta) (pathdb.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(pathdb.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockReadWriteMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockReadWrite)(nil).Insert), arg0, arg1)
}

// InsertNextQuery mocks base method
func (m *MockReadWrite) InsertNextQuery(arg0 context.Context, arg1, arg2 addr.IA, arg3 pathdb.PolicyHash, arg4 time.Time) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNextQuery", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertNextQuery indicates an expected call of InsertNextQuery
func (mr *MockReadWriteMockRecorder) InsertNextQuery(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNextQuery", reflect.TypeOf((*MockReadWrite)(nil).InsertNextQuery), arg0, arg1, arg2, arg3, arg4)
}

// InsertWithHPCfgIDs mocks base method
func (m *MockReadWrite) InsertWithHPCfgIDs(arg0 context.Context, arg1 *seg.Meta, arg2 []*query.HPCfgID) (pathdb.InsertStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertWithHPCfgIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].(pathdb.InsertStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertWithHPCfgIDs indicates an expected call of InsertWithHPCfgIDs
func (mr *MockReadWriteMockRecorder) InsertWithHPCfgIDs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertWithHPCfgIDs", reflect.TypeOf((*MockReadWrite)(nil).InsertWithHPCfgIDs), arg0, arg1, arg2)
}