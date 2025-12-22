package wifi_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
	myWifi "github.com/verticalochka/task-6/internal/wifi"
)

var (
	errTestInterfaces = fmt.Errorf("test interfaces error")
)

func createMAC(s string) net.HardwareAddr {
	addr, _ := net.ParseMAC(s)

	return addr
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := myWifi.New(mockHandler)

	networkIfaces := []*wifi.Interface{
		{Name: "wifi0", HardwareAddr: createMAC("11:22:33:44:55:66")},
		{Name: "wifi1", HardwareAddr: createMAC("aa:bb:cc:dd:ee:ff")},
	}

	mockHandler.On("Interfaces").Return(networkIfaces, nil)

	macs, err := service.GetAddresses()
	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{
		createMAC("11:22:33:44:55:66"),
		createMAC("aa:bb:cc:dd:ee:ff"),
	}, macs)
}

func TestGetAddresses_Failed(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := myWifi.New(mockHandler)

	mockHandler.On("Interfaces").Return(nil, errTestInterfaces)

	macs, err := service.GetAddresses()
	require.ErrorContains(t, err, "getting interfaces")
	require.Nil(t, macs)
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := myWifi.New(mockHandler)

	networkIfaces := []*wifi.Interface{
		{Name: "wireless0", HardwareAddr: createMAC("11:22:33:44:55:66")},
		{Name: "ethernet1", HardwareAddr: createMAC("aa:bb:cc:dd:ee:ff")},
	}

	mockHandler.On("Interfaces").Return(networkIfaces, nil)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"wireless0", "ethernet1"}, names)
}

func TestGetNames_Failed(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := myWifi.New(mockHandler)

	mockHandler.On("Interfaces").Return(nil, errTestInterfaces)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "getting interfaces")
	require.Nil(t, names)
}

func TestGetAddresses_EmptyResult(t *testing.T) {
	t.Parallel()

	mockHandler := NewWiFiHandle(t)
	service := myWifi.New(mockHandler)

	mockHandler.On("Interfaces").Return([]*wifi.Interface{}, nil)

	macs, err := service.GetAddresses()
	require.NoError(t, err)
	require.Empty(t, macs)
}
