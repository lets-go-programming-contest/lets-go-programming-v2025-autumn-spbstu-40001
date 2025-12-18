package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/DimasFantomasA/task-6/internal/wifi"

	mdwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

func TestWiFiService_GetAddresses(t *testing.T) {
	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	tests := []struct {
		name        string
		ifaces      []*mdwifi.Interface
		errExpected error
		expected    []net.HardwareAddr
	}{
		{
			name: "success",
			ifaces: []*mdwifi.Interface{
				{HardwareAddr: mustMAC("00:11:22:33:44:55")},
				{HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
			},
			expected: []net.HardwareAddr{
				mustMAC("00:11:22:33:44:55"),
				mustMAC("aa:bb:cc:dd:ee:ff"),
			},
		},
		{
			name:        "error",
			errExpected: errors.New("wifi error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWiFi.On("Interfaces").Return(tt.ifaces, tt.errExpected)

			result, err := service.GetAddresses()

			if tt.errExpected != nil {
				require.ErrorIs(t, err, tt.errExpected)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestWiFiService_GetNames(t *testing.T) {
	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*mdwifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
	}, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "eth0"}, names)
}

func mustMAC(s string) net.HardwareAddr {
	mac, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}
	return mac
}
