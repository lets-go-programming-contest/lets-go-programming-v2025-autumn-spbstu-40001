package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	mywifi "github.com/A1exMas1ov/task-6/internal/wifi"
)

var errWiFi = errors.New("wifi error")

func mac(s string) net.HardwareAddr {
	addr, _ := net.ParseMAC(s)
	return addr
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		ifaces     []*wifi.Interface
		err        error
		expectErr  bool
		expectAddr []net.HardwareAddr
	}{
		{
			name: "success",
			ifaces: []*wifi.Interface{
				{HardwareAddr: mac("12:34:56:78:9a:bc")},
				{HardwareAddr: mac("de:ad:be:ef:00:01")},
			},
			expectAddr: []net.HardwareAddr{
				mac("12:34:56:78:9a:bc"),
				mac("de:ad:be:ef:00:01"),
			},
		},
		{
			name:      "interfaces error",
			err:       errWiFi,
			expectErr: true,
		},
		{
			name:       "empty interfaces",
			ifaces:     []*wifi.Interface{},
			expectAddr: []net.HardwareAddr{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockWiFi := NewWiFiHandle(t)
			service := mywifi.New(mockWiFi)

			mockWiFi.On("Interfaces").Return(tt.ifaces, tt.err)

			addrs, err := service.GetAddresses()

			if tt.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectAddr, addrs)
			}
		})
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	service := mywifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
	}, nil)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "eth0"}, names)
}
