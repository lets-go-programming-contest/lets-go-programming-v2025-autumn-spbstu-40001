package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	myWifi "github.com/A1exCRE/task-6/internal/wifi"
)

var testErr = errors.New("test error")

func TestWiFiNew(t *testing.T) {
	mockHandle := NewWiFiHandle(t)
	service := myWifi.New(mockHandle)
	assert.Equal(t, mockHandle, service.WiFi)
}

func TestGetAddressesBasic(t *testing.T) {
	mockHandle := NewWiFiHandle(t)
	service := myWifi.WiFiService{WiFi: mockHandle}

	hwAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	ifaces := []*wifi.Interface{
		{
			Index:        1,
			Name:         "wlan0",
			HardwareAddr: hwAddr1,
		},
	}

	mockHandle.On("Interfaces").Return(ifaces, nil)

	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	require.Len(t, addrs, 1)
	require.Equal(t, hwAddr1, addrs[0])
	mockHandle.AssertExpectations(t)
}

func TestGetNamesBasic(t *testing.T) {
	mockHandle := NewWiFiHandle(t)
	service := myWifi.WiFiService{WiFi: mockHandle}

	ifaces := []*wifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
	}

	mockHandle.On("Interfaces").Return(ifaces, nil)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "eth0"}, names)
	mockHandle.AssertExpectations(t)
}

func TestGetAddressesError(t *testing.T) {
	mockHandle := NewWiFiHandle(t)
	service := myWifi.WiFiService{WiFi: mockHandle}

	mockHandle.On("Interfaces").Return(nil, testErr)

	addrs, err := service.GetAddresses()

	require.Error(t, err)
	require.ErrorContains(t, err, "getting interfaces")
	require.Nil(t, addrs)
	mockHandle.AssertExpectations(t)
}
