package wifi

import (
	"fmt"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
)

type WiFiHandle struct {
	mock.Mock
}

func (m *WiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()

	ifaces, _ := args.Get(0).([]*wifi.Interface)

	err := args.Error(1)
	if err != nil {
		err = fmt.Errorf("mock error: %w", err)
	}

	return ifaces, err
}
