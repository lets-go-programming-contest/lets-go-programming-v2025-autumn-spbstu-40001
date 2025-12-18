package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/DimasFantomasA/task-6/internal/wifi"

	mdwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	require.NotNil(t, service)
	require.NotNil(t, service.WiFi)
}

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
			name:        "error from Interfaces",
			errExpected: errors.New("wifi error"),
		},
		{
			name:     "empty result",
			ifaces:   []*mdwifi.Interface{},
			expected: []net.HardwareAddr{},
		},
		{
			name: "interface with nil hardware address",
			ifaces: []*mdwifi.Interface{
				{HardwareAddr: nil},
				{HardwareAddr: mustMAC("00:11:22:33:44:55")},
			},
			expected: []net.HardwareAddr{
				nil,
				mustMAC("00:11:22:33:44:55"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWiFi.On("Interfaces").Return(tt.ifaces, tt.errExpected)

			result, err := service.GetAddresses()

			if tt.errExpected != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), "getting interfaces")
				require.Contains(t, err.Error(), tt.errExpected.Error())
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

	tests := []struct {
		name        string
		ifaces      []*mdwifi.Interface
		errExpected error
		expected    []string
	}{
		{
			name: "success",
			ifaces: []*mdwifi.Interface{
				{Name: "wlan0"},
				{Name: "eth0"},
				{Name: "wlan1"},
			},
			expected: []string{"wlan0", "eth0", "wlan1"},
		},
		{
			name:        "error from Interfaces",
			errExpected: errors.New("wifi error"),
		},
		{
			name:     "empty result",
			ifaces:   []*mdwifi.Interface{},
			expected: []string{},
		},
		{
			name: "interface with empty name",
			ifaces: []*mdwifi.Interface{
				{Name: ""},
				{Name: "wlan0"},
			},
			expected: []string{"", "wlan0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWiFi.On("Interfaces").Return(tt.ifaces, tt.errExpected)

			result, err := service.GetNames()

			if tt.errExpected != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), "getting interfaces")
				require.Contains(t, err.Error(), tt.errExpected.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func mustMAC(s string) net.HardwareAddr {
	mac, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}
	return mac
}
