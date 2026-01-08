package wifi_test

import (
	"errors"
	"net"
	"testing"

	myWifi "github.com/atroxxxxxx/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var errWifi = errors.New("wifi")

type testT struct {
}

func hardwareAddr(str string) net.HardwareAddr {
	addr, _ := net.ParseMAC(str)
	return addr
}

func TestWiFiService_GetAddresses(test *testing.T) {
	test.Parallel()

	tests := []struct {
		name          string
		interfaces    []*wifi.Interface
		err           error
		expectedError bool
		expectedAddr  []net.HardwareAddr
	}{
		{
			name: "success",
			interfaces: []*wifi.Interface{
				{HardwareAddr: hardwareAddr("12:34:56:78:9a:bc")},
				{HardwareAddr: hardwareAddr("ab:cd:de:ef:00:01")},
			},
			expectedAddr: []net.HardwareAddr{
				hardwareAddr("12:34:56:78:9a:bc"),
				hardwareAddr("ab:cd:de:ef:00:01"),
			},
		},
		{
			name:         "empty",
			interfaces:   []*wifi.Interface{},
			expectedAddr: []net.HardwareAddr{},
		},

		{
			name:          "interface err",
			err:           errWifi,
			expectedError: true,
		},
	}
	for _, t := range tests {
		test.Run(t.name, func(t2 *testing.T) {
			t2.Parallel()

			mock := NewWiFiHandle(t2)
			service := myWifi.New(mock)

			mock.On("Interfaces").Return(t.interfaces, t.err)
			addresses, err := service.GetAddresses()

			if t.expectedError {
				require.Error(t2, err)
			} else {
				require.NoError(t2, err)
				require.Equal(t2, t.expectedAddr, addresses)
			}
		})
	}
}

func TestWiFiService_GetNames(test *testing.T) {
	test.Parallel()

	tests := []struct {
		name        string
		interfaces  []*wifi.Interface
		err         error
		expectError bool
		expect      []string
	}{
		{
			name: "success",
			interfaces: []*wifi.Interface{
				{Name: "lan1"},
				{Name: "lan2"},
				{Name: "wifi"},
			},
			expect: []string{"lan1", "lan2", "wifi"},
		},
		{
			name:       "empty",
			interfaces: []*wifi.Interface{},
			expect:     []string{},
		},

		{
			name:        "interface error",
			err:         errWifi,
			expectError: true,
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(t2 *testing.T) {
			t2.Parallel()

			mock := NewWiFiHandle(t2)
			service := myWifi.New(mock)

			mock.On("Interfaces").Return(t.interfaces, t.err)

			names, err := service.GetNames()

			if t.expectError {
				require.Error(t2, err)
			} else {
				require.NoError(t2, err)
				require.Equal(t2, t.expect, names)
			}
		})
	}
}
