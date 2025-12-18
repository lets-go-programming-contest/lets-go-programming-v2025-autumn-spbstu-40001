package wifi_test

import (
	"fmt"

	wifi "github.com/mdlayher/wifi"
	mock "github.com/stretchr/testify/mock"
)

type WiFiHandle struct {
	mock.Mock
}

func (_m *WiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Interfaces")
	}

	var (
		r0 []*wifi.Interface
		r1 error
	)

	if rf, ok := ret.Get(0).(func() ([]*wifi.Interface, error)); ok {
		return rf()
	}

	if val := ret.Get(0); val != nil {
		if rf, ok := val.(func() []*wifi.Interface); ok {
			r0 = rf()
		} else if arr, ok := val.([]*wifi.Interface); ok {
			r0 = arr
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		if err := ret.Error(1); err != nil {
			r1 = fmt.Errorf("mock error: %w", err)
		}
	}

	return r0, r1
}

func NewWiFiHandle(t interface {
	mock.TestingT
	Cleanup(cleanupFunc func())
},
) *WiFiHandle {
	mock := &WiFiHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
