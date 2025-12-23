package wifi_test

import (
	"fmt"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
)

type WiFiHandleMock struct {
	mock.Mock
}

func (m *WiFiHandleMock) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()

	var wifiInterfaces []*wifi.Interface

	if data, ok := args.Get(0).(func() []*wifi.Interface); ok {
		wifiInterfaces = data()
	} else if data, ok := args.Get(0).([]*wifi.Interface); ok {
		wifiInterfaces = data
	}

	var executionError error

	if errData := args.Error(1); errData != nil {
		executionError = fmt.Errorf("mock error: %w", errData)
	}

	return wifiInterfaces, executionError
}
