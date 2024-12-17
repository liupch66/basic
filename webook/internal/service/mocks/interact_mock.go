// Code generated by MockGen. DO NOT EDIT.
// Source: interact.go
//
// Generated by this command:
//
//	mockgen -package=svcmocks -source=interact.go -destination=mocks/interact_mock.go InteractService
//

// Package svcmocks is a generated GoMock package.
package svcmocks

import (
	domain "basic-go/webook/internal/domain"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInteractService is a mock of InteractService interface.
type MockInteractService struct {
	ctrl     *gomock.Controller
	recorder *MockInteractServiceMockRecorder
	isgomock struct{}
}

// MockInteractServiceMockRecorder is the mock recorder for MockInteractService.
type MockInteractServiceMockRecorder struct {
	mock *MockInteractService
}

// NewMockInteractService creates a new mock instance.
func NewMockInteractService(ctrl *gomock.Controller) *MockInteractService {
	mock := &MockInteractService{ctrl: ctrl}
	mock.recorder = &MockInteractServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractService) EXPECT() *MockInteractServiceMockRecorder {
	return m.recorder
}

// CancelLike mocks base method.
func (m *MockInteractService) CancelLike(ctx context.Context, biz string, bizId, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelLike", ctx, biz, bizId, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelLike indicates an expected call of CancelLike.
func (mr *MockInteractServiceMockRecorder) CancelLike(ctx, biz, bizId, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelLike", reflect.TypeOf((*MockInteractService)(nil).CancelLike), ctx, biz, bizId, uid)
}

// Collect mocks base method.
func (m *MockInteractService) Collect(ctx context.Context, biz string, bizId, cid, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collect", ctx, biz, bizId, cid, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Collect indicates an expected call of Collect.
func (mr *MockInteractServiceMockRecorder) Collect(ctx, biz, bizId, cid, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockInteractService)(nil).Collect), ctx, biz, bizId, cid, uid)
}

// Get mocks base method.
func (m *MockInteractService) Get(ctx context.Context, biz string, bizId, uid int64) (domain.Interact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, biz, bizId, uid)
	ret0, _ := ret[0].(domain.Interact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInteractServiceMockRecorder) Get(ctx, biz, bizId, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInteractService)(nil).Get), ctx, biz, bizId, uid)
}

// GetByIds mocks base method.
func (m *MockInteractService) GetByIds(ctx context.Context, biz string, ids []int64) (map[int64]domain.Interact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIds", ctx, biz, ids)
	ret0, _ := ret[0].(map[int64]domain.Interact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIds indicates an expected call of GetByIds.
func (mr *MockInteractServiceMockRecorder) GetByIds(ctx, biz, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIds", reflect.TypeOf((*MockInteractService)(nil).GetByIds), ctx, biz, ids)
}

// IncrReadCnt mocks base method.
func (m *MockInteractService) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrReadCnt", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrReadCnt indicates an expected call of IncrReadCnt.
func (mr *MockInteractServiceMockRecorder) IncrReadCnt(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrReadCnt", reflect.TypeOf((*MockInteractService)(nil).IncrReadCnt), ctx, biz, bizId)
}

// Like mocks base method.
func (m *MockInteractService) Like(ctx context.Context, biz string, id, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Like", ctx, biz, id, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Like indicates an expected call of Like.
func (mr *MockInteractServiceMockRecorder) Like(ctx, biz, id, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Like", reflect.TypeOf((*MockInteractService)(nil).Like), ctx, biz, id, uid)
}
