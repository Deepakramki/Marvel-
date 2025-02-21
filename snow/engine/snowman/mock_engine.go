// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/snow/engine/snowman (interfaces: Engine)

// Package snowman is a generated GoMock package.
package snowman

import (
	context "context"
	reflect "reflect"
	time "time"

	ids "github.com/ava-labs/avalanchego/ids"
	snow "github.com/ava-labs/avalanchego/snow"
	snowman "github.com/ava-labs/avalanchego/snow/consensus/snowman"
	common "github.com/ava-labs/avalanchego/snow/engine/common"
	version "github.com/ava-labs/avalanchego/version"
	gomock "github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Accepted mocks base method.
func (m *MockEngine) Accepted(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accepted", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Accepted indicates an expected call of Accepted.
func (mr *MockEngineMockRecorder) Accepted(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accepted", reflect.TypeOf((*MockEngine)(nil).Accepted), arg0, arg1, arg2, arg3)
}

// AcceptedFrontier mocks base method.
func (m *MockEngine) AcceptedFrontier(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedFrontier", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptedFrontier indicates an expected call of AcceptedFrontier.
func (mr *MockEngineMockRecorder) AcceptedFrontier(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedFrontier", reflect.TypeOf((*MockEngine)(nil).AcceptedFrontier), arg0, arg1, arg2, arg3)
}

// AcceptedStateSummary mocks base method.
func (m *MockEngine) AcceptedStateSummary(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptedStateSummary", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptedStateSummary indicates an expected call of AcceptedStateSummary.
func (mr *MockEngineMockRecorder) AcceptedStateSummary(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptedStateSummary", reflect.TypeOf((*MockEngine)(nil).AcceptedStateSummary), arg0, arg1, arg2, arg3)
}

// Ancestors mocks base method.
func (m *MockEngine) Ancestors(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 [][]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ancestors", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ancestors indicates an expected call of Ancestors.
func (mr *MockEngineMockRecorder) Ancestors(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ancestors", reflect.TypeOf((*MockEngine)(nil).Ancestors), arg0, arg1, arg2, arg3)
}

// AppGossip mocks base method.
func (m *MockEngine) AppGossip(arg0 context.Context, arg1 ids.NodeID, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppGossip", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppGossip indicates an expected call of AppGossip.
func (mr *MockEngineMockRecorder) AppGossip(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppGossip", reflect.TypeOf((*MockEngine)(nil).AppGossip), arg0, arg1, arg2)
}

// AppRequest mocks base method.
func (m *MockEngine) AppRequest(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 time.Time, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequest indicates an expected call of AppRequest.
func (mr *MockEngineMockRecorder) AppRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequest", reflect.TypeOf((*MockEngine)(nil).AppRequest), arg0, arg1, arg2, arg3, arg4)
}

// AppRequestFailed mocks base method.
func (m *MockEngine) AppRequestFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppRequestFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppRequestFailed indicates an expected call of AppRequestFailed.
func (mr *MockEngineMockRecorder) AppRequestFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppRequestFailed", reflect.TypeOf((*MockEngine)(nil).AppRequestFailed), arg0, arg1, arg2)
}

// AppResponse mocks base method.
func (m *MockEngine) AppResponse(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppResponse indicates an expected call of AppResponse.
func (mr *MockEngineMockRecorder) AppResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppResponse", reflect.TypeOf((*MockEngine)(nil).AppResponse), arg0, arg1, arg2, arg3)
}

// Chits mocks base method.
func (m *MockEngine) Chits(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chits", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chits indicates an expected call of Chits.
func (mr *MockEngineMockRecorder) Chits(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chits", reflect.TypeOf((*MockEngine)(nil).Chits), arg0, arg1, arg2, arg3)
}

// Connected mocks base method.
func (m *MockEngine) Connected(arg0 ids.NodeID, arg1 *version.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connected indicates an expected call of Connected.
func (mr *MockEngineMockRecorder) Connected(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockEngine)(nil).Connected), arg0, arg1)
}

// Context mocks base method.
func (m *MockEngine) Context() *snow.ConsensusContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(*snow.ConsensusContext)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockEngineMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEngine)(nil).Context))
}

// CrossChainAppRequest mocks base method.
func (m *MockEngine) CrossChainAppRequest(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 time.Time, arg4 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppRequest indicates an expected call of CrossChainAppRequest.
func (mr *MockEngineMockRecorder) CrossChainAppRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppRequest", reflect.TypeOf((*MockEngine)(nil).CrossChainAppRequest), arg0, arg1, arg2, arg3, arg4)
}

// CrossChainAppRequestFailed mocks base method.
func (m *MockEngine) CrossChainAppRequestFailed(arg0 context.Context, arg1 ids.ID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppRequestFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppRequestFailed indicates an expected call of CrossChainAppRequestFailed.
func (mr *MockEngineMockRecorder) CrossChainAppRequestFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppRequestFailed", reflect.TypeOf((*MockEngine)(nil).CrossChainAppRequestFailed), arg0, arg1, arg2)
}

// CrossChainAppResponse mocks base method.
func (m *MockEngine) CrossChainAppResponse(arg0 context.Context, arg1 ids.ID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrossChainAppResponse", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CrossChainAppResponse indicates an expected call of CrossChainAppResponse.
func (mr *MockEngineMockRecorder) CrossChainAppResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrossChainAppResponse", reflect.TypeOf((*MockEngine)(nil).CrossChainAppResponse), arg0, arg1, arg2, arg3)
}

// Disconnected mocks base method.
func (m *MockEngine) Disconnected(arg0 ids.NodeID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnected", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnected indicates an expected call of Disconnected.
func (mr *MockEngineMockRecorder) Disconnected(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnected", reflect.TypeOf((*MockEngine)(nil).Disconnected), arg0)
}

// Get mocks base method.
func (m *MockEngine) Get(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockEngineMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockEngine)(nil).Get), arg0, arg1, arg2, arg3)
}

// GetAccepted mocks base method.
func (m *MockEngine) GetAccepted(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccepted", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAccepted indicates an expected call of GetAccepted.
func (mr *MockEngineMockRecorder) GetAccepted(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccepted", reflect.TypeOf((*MockEngine)(nil).GetAccepted), arg0, arg1, arg2, arg3)
}

// GetAcceptedFailed mocks base method.
func (m *MockEngine) GetAcceptedFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAcceptedFailed indicates an expected call of GetAcceptedFailed.
func (mr *MockEngineMockRecorder) GetAcceptedFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedFailed", reflect.TypeOf((*MockEngine)(nil).GetAcceptedFailed), arg0, arg1, arg2)
}

// GetAcceptedFrontier mocks base method.
func (m *MockEngine) GetAcceptedFrontier(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAcceptedFrontier indicates an expected call of GetAcceptedFrontier.
func (mr *MockEngineMockRecorder) GetAcceptedFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedFrontier", reflect.TypeOf((*MockEngine)(nil).GetAcceptedFrontier), arg0, arg1, arg2)
}

// GetAcceptedFrontierFailed mocks base method.
func (m *MockEngine) GetAcceptedFrontierFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedFrontierFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAcceptedFrontierFailed indicates an expected call of GetAcceptedFrontierFailed.
func (mr *MockEngineMockRecorder) GetAcceptedFrontierFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedFrontierFailed", reflect.TypeOf((*MockEngine)(nil).GetAcceptedFrontierFailed), arg0, arg1, arg2)
}

// GetAcceptedStateSummary mocks base method.
func (m *MockEngine) GetAcceptedStateSummary(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedStateSummary", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAcceptedStateSummary indicates an expected call of GetAcceptedStateSummary.
func (mr *MockEngineMockRecorder) GetAcceptedStateSummary(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedStateSummary", reflect.TypeOf((*MockEngine)(nil).GetAcceptedStateSummary), arg0, arg1, arg2, arg3)
}

// GetAcceptedStateSummaryFailed mocks base method.
func (m *MockEngine) GetAcceptedStateSummaryFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAcceptedStateSummaryFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAcceptedStateSummaryFailed indicates an expected call of GetAcceptedStateSummaryFailed.
func (mr *MockEngineMockRecorder) GetAcceptedStateSummaryFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAcceptedStateSummaryFailed", reflect.TypeOf((*MockEngine)(nil).GetAcceptedStateSummaryFailed), arg0, arg1, arg2)
}

// GetAncestors mocks base method.
func (m *MockEngine) GetAncestors(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAncestors", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAncestors indicates an expected call of GetAncestors.
func (mr *MockEngineMockRecorder) GetAncestors(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAncestors", reflect.TypeOf((*MockEngine)(nil).GetAncestors), arg0, arg1, arg2, arg3)
}

// GetAncestorsFailed mocks base method.
func (m *MockEngine) GetAncestorsFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAncestorsFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAncestorsFailed indicates an expected call of GetAncestorsFailed.
func (mr *MockEngineMockRecorder) GetAncestorsFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAncestorsFailed", reflect.TypeOf((*MockEngine)(nil).GetAncestorsFailed), arg0, arg1, arg2)
}

// GetBlock mocks base method.
func (m *MockEngine) GetBlock(arg0 ids.ID) (snowman.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlock", arg0)
	ret0, _ := ret[0].(snowman.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlock indicates an expected call of GetBlock.
func (mr *MockEngineMockRecorder) GetBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlock", reflect.TypeOf((*MockEngine)(nil).GetBlock), arg0)
}

// GetFailed mocks base method.
func (m *MockEngine) GetFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetFailed indicates an expected call of GetFailed.
func (mr *MockEngineMockRecorder) GetFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFailed", reflect.TypeOf((*MockEngine)(nil).GetFailed), arg0, arg1, arg2)
}

// GetStateSummaryFrontier mocks base method.
func (m *MockEngine) GetStateSummaryFrontier(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateSummaryFrontier", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetStateSummaryFrontier indicates an expected call of GetStateSummaryFrontier.
func (mr *MockEngineMockRecorder) GetStateSummaryFrontier(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateSummaryFrontier", reflect.TypeOf((*MockEngine)(nil).GetStateSummaryFrontier), arg0, arg1, arg2)
}

// GetStateSummaryFrontierFailed mocks base method.
func (m *MockEngine) GetStateSummaryFrontierFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateSummaryFrontierFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetStateSummaryFrontierFailed indicates an expected call of GetStateSummaryFrontierFailed.
func (mr *MockEngineMockRecorder) GetStateSummaryFrontierFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateSummaryFrontierFailed", reflect.TypeOf((*MockEngine)(nil).GetStateSummaryFrontierFailed), arg0, arg1, arg2)
}

// GetVM mocks base method.
func (m *MockEngine) GetVM() common.VM {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVM")
	ret0, _ := ret[0].(common.VM)
	return ret0
}

// GetVM indicates an expected call of GetVM.
func (mr *MockEngineMockRecorder) GetVM() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVM", reflect.TypeOf((*MockEngine)(nil).GetVM))
}

// Gossip mocks base method.
func (m *MockEngine) Gossip() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Gossip")
	ret0, _ := ret[0].(error)
	return ret0
}

// Gossip indicates an expected call of Gossip.
func (mr *MockEngineMockRecorder) Gossip() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gossip", reflect.TypeOf((*MockEngine)(nil).Gossip))
}

// Halt mocks base method.
func (m *MockEngine) Halt() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Halt")
}

// Halt indicates an expected call of Halt.
func (mr *MockEngineMockRecorder) Halt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Halt", reflect.TypeOf((*MockEngine)(nil).Halt))
}

// HealthCheck mocks base method.
func (m *MockEngine) HealthCheck() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockEngineMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockEngine)(nil).HealthCheck))
}

// Notify mocks base method.
func (m *MockEngine) Notify(arg0 common.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Notify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Notify indicates an expected call of Notify.
func (mr *MockEngineMockRecorder) Notify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockEngine)(nil).Notify), arg0)
}

// PullQuery mocks base method.
func (m *MockEngine) PullQuery(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 ids.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullQuery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullQuery indicates an expected call of PullQuery.
func (mr *MockEngineMockRecorder) PullQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullQuery", reflect.TypeOf((*MockEngine)(nil).PullQuery), arg0, arg1, arg2, arg3)
}

// PushQuery mocks base method.
func (m *MockEngine) PushQuery(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushQuery", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushQuery indicates an expected call of PushQuery.
func (mr *MockEngineMockRecorder) PushQuery(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushQuery", reflect.TypeOf((*MockEngine)(nil).PushQuery), arg0, arg1, arg2, arg3)
}

// Put mocks base method.
func (m *MockEngine) Put(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockEngineMockRecorder) Put(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockEngine)(nil).Put), arg0, arg1, arg2, arg3)
}

// QueryFailed mocks base method.
func (m *MockEngine) QueryFailed(arg0 context.Context, arg1 ids.NodeID, arg2 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFailed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryFailed indicates an expected call of QueryFailed.
func (mr *MockEngineMockRecorder) QueryFailed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFailed", reflect.TypeOf((*MockEngine)(nil).QueryFailed), arg0, arg1, arg2)
}

// Shutdown mocks base method.
func (m *MockEngine) Shutdown() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown")
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockEngineMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockEngine)(nil).Shutdown))
}

// Start mocks base method.
func (m *MockEngine) Start(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockEngineMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEngine)(nil).Start), arg0)
}

// StateSummaryFrontier mocks base method.
func (m *MockEngine) StateSummaryFrontier(arg0 context.Context, arg1 ids.NodeID, arg2 uint32, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateSummaryFrontier", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// StateSummaryFrontier indicates an expected call of StateSummaryFrontier.
func (mr *MockEngineMockRecorder) StateSummaryFrontier(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateSummaryFrontier", reflect.TypeOf((*MockEngine)(nil).StateSummaryFrontier), arg0, arg1, arg2, arg3)
}

// Timeout mocks base method.
func (m *MockEngine) Timeout() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Timeout indicates an expected call of Timeout.
func (mr *MockEngineMockRecorder) Timeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockEngine)(nil).Timeout))
}
