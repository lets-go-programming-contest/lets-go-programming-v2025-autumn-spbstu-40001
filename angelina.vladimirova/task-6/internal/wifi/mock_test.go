package wifi_test

import (
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
)

type mockWiFiHandle struct {
	mock.Mock
}

func (m *mockWiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()
	return args.Get(0).([]*wifi.Interface), args.Error(1)
}

func (m *mockWiFiHandle) StationInfo(ifi *wifi.Interface) (*wifi.StationInfo, error) {
	args := m.Called(ifi)
	return args.Get(0).(*wifi.StationInfo), args.Error(1)
}

func newMockWiFiHandle(t mock.TestingT) *mockWiFiHandle {
	mock := &mockWiFiHandle{}
	mock.Test(t)
	return mock
}
