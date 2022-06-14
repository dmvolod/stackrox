// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	compliance "github.com/stackrox/rox/central/compliance"
	types "github.com/stackrox/rox/central/compliance/datastore/types"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// ClearAggregationResults mocks base method.
func (m *MockStore) ClearAggregationResults() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearAggregationResults")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearAggregationResults indicates an expected call of ClearAggregationResults.
func (mr *MockStoreMockRecorder) ClearAggregationResults() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearAggregationResults", reflect.TypeOf((*MockStore)(nil).ClearAggregationResults))
}

// GetAggregationResult mocks base method.
func (m *MockStore) GetAggregationResult(queryString string, groupBy []storage.ComplianceAggregation_Scope, unit storage.ComplianceAggregation_Scope) ([]*storage.ComplianceAggregation_Result, []*storage.ComplianceAggregation_Source, map[*storage.ComplianceAggregation_Result]*storage.ComplianceDomain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAggregationResult", queryString, groupBy, unit)
	ret0, _ := ret[0].([]*storage.ComplianceAggregation_Result)
	ret1, _ := ret[1].([]*storage.ComplianceAggregation_Source)
	ret2, _ := ret[2].(map[*storage.ComplianceAggregation_Result]*storage.ComplianceDomain)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetAggregationResult indicates an expected call of GetAggregationResult.
func (mr *MockStoreMockRecorder) GetAggregationResult(queryString, groupBy, unit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAggregationResult", reflect.TypeOf((*MockStore)(nil).GetAggregationResult), queryString, groupBy, unit)
}

// GetLatestRunMetadataBatch mocks base method.
func (m *MockStore) GetLatestRunMetadataBatch(clusterID string, standardIDs []string) (map[compliance.ClusterStandardPair]types.ComplianceRunsMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRunMetadataBatch", clusterID, standardIDs)
	ret0, _ := ret[0].(map[compliance.ClusterStandardPair]types.ComplianceRunsMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRunMetadataBatch indicates an expected call of GetLatestRunMetadataBatch.
func (mr *MockStoreMockRecorder) GetLatestRunMetadataBatch(clusterID, standardIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRunMetadataBatch", reflect.TypeOf((*MockStore)(nil).GetLatestRunMetadataBatch), clusterID, standardIDs)
}

// GetLatestRunResults mocks base method.
func (m *MockStore) GetLatestRunResults(clusterID, standardID string, flags types.GetFlags) (types.ResultsWithStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRunResults", clusterID, standardID, flags)
	ret0, _ := ret[0].(types.ResultsWithStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRunResults indicates an expected call of GetLatestRunResults.
func (mr *MockStoreMockRecorder) GetLatestRunResults(clusterID, standardID, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRunResults", reflect.TypeOf((*MockStore)(nil).GetLatestRunResults), clusterID, standardID, flags)
}

// GetLatestRunResultsBatch mocks base method.
func (m *MockStore) GetLatestRunResultsBatch(clusterIDs, standardIDs []string, flags types.GetFlags) (map[compliance.ClusterStandardPair]types.ResultsWithStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestRunResultsBatch", clusterIDs, standardIDs, flags)
	ret0, _ := ret[0].(map[compliance.ClusterStandardPair]types.ResultsWithStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestRunResultsBatch indicates an expected call of GetLatestRunResultsBatch.
func (mr *MockStoreMockRecorder) GetLatestRunResultsBatch(clusterIDs, standardIDs, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestRunResultsBatch", reflect.TypeOf((*MockStore)(nil).GetLatestRunResultsBatch), clusterIDs, standardIDs, flags)
}

// GetSpecificRunResults mocks base method.
func (m *MockStore) GetSpecificRunResults(clusterID, standardID, runID string, flags types.GetFlags) (types.ResultsWithStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpecificRunResults", clusterID, standardID, runID, flags)
	ret0, _ := ret[0].(types.ResultsWithStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpecificRunResults indicates an expected call of GetSpecificRunResults.
func (mr *MockStoreMockRecorder) GetSpecificRunResults(clusterID, standardID, runID, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpecificRunResults", reflect.TypeOf((*MockStore)(nil).GetSpecificRunResults), clusterID, standardID, runID, flags)
}

// StoreAggregationResult mocks base method.
func (m *MockStore) StoreAggregationResult(queryString string, groupBy []storage.ComplianceAggregation_Scope, unit storage.ComplianceAggregation_Scope, results []*storage.ComplianceAggregation_Result, sources []*storage.ComplianceAggregation_Source, domains map[*storage.ComplianceAggregation_Result]*storage.ComplianceDomain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAggregationResult", queryString, groupBy, unit, results, sources, domains)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreAggregationResult indicates an expected call of StoreAggregationResult.
func (mr *MockStoreMockRecorder) StoreAggregationResult(queryString, groupBy, unit, results, sources, domains interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAggregationResult", reflect.TypeOf((*MockStore)(nil).StoreAggregationResult), queryString, groupBy, unit, results, sources, domains)
}

// StoreComplianceDomain mocks base method.
func (m *MockStore) StoreComplianceDomain(domain *storage.ComplianceDomain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreComplianceDomain", domain)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreComplianceDomain indicates an expected call of StoreComplianceDomain.
func (mr *MockStoreMockRecorder) StoreComplianceDomain(domain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreComplianceDomain", reflect.TypeOf((*MockStore)(nil).StoreComplianceDomain), domain)
}

// StoreFailure mocks base method.
func (m *MockStore) StoreFailure(metadata *storage.ComplianceRunMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreFailure", metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreFailure indicates an expected call of StoreFailure.
func (mr *MockStoreMockRecorder) StoreFailure(metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreFailure", reflect.TypeOf((*MockStore)(nil).StoreFailure), metadata)
}

// StoreRunResults mocks base method.
func (m *MockStore) StoreRunResults(results *storage.ComplianceRunResults) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreRunResults", results)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreRunResults indicates an expected call of StoreRunResults.
func (mr *MockStoreMockRecorder) StoreRunResults(results interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreRunResults", reflect.TypeOf((*MockStore)(nil).StoreRunResults), results)
}
