package wifi_test

import (
	"errors"
	"net"
	"testing"

	taskWifiPack "github.com/Aapng-cmd/task-6/internal/wifi"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ErrExpected = errors.New("expected error")

func TestWiFiServiceGetAddressesSuccess(t *testing.T) {
	t.Parallel()

	mockWiFi := new(WiFiHandle)

	hwAddr1, _ := net.ParseMAC("00:11:22:33:44:55")
	hwAddr2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")

	numberOfData := 3

	interfaces := []*wifi.Interface{
		{
			Name: "wlan0", HardwareAddr: hwAddr1,
		},
		{
			Name: "wlan1", HardwareAddr: hwAddr2,
		},
		{
			Name: "wlan2", HardwareAddr: hwAddr2,
		},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	service := taskWifiPack.New(mockWiFi)
	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Len(t, addrs, numberOfData)
	assert.Equal(t, hwAddr1, addrs[0])
	assert.Equal(t, hwAddr2, addrs[1])
	assert.Equal(t, hwAddr2, addrs[2])

	mockWiFi.AssertExpectations(t)
}

func TestWiFiServiceGetAddressesError(t *testing.T) {
	t.Parallel()

	mockWiFi := new(WiFiHandle)
	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, ErrExpected)

	service := taskWifiPack.New(mockWiFi)
	addrs, err := service.GetAddresses()

	require.Error(t, err)
	assert.Nil(t, addrs)
	assert.Contains(t, err.Error(), "getting interfaces")

	mockWiFi.AssertExpectations(t)
}

func TestWiFiServiceGetNamesSuccess(t *testing.T) {
	const numberOfData = 3

	t.Parallel()

	mockWiFi := new(WiFiHandle)

	hwAddr, _ := net.ParseMAC("13:37:de:ad:be:ef")
	interfaces := []*wifi.Interface{
		{
			Name: "wlan0", HardwareAddr: hwAddr,
		},
		{Name: "wlan1"},
		{Name: "eth0"},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	service := taskWifiPack.New(mockWiFi)
	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Len(t, names, numberOfData)
	assert.Equal(t, []string{"wlan0", "wlan1", "eth0"}, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiServiceGetNamesEmpty(t *testing.T) {
	t.Parallel()

	mockWiFi := new(WiFiHandle)
	interfaces := []*wifi.Interface{}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	service := taskWifiPack.New(mockWiFi)
	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Empty(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiServiceGetNamesError(t *testing.T) {
	t.Parallel()

	mockWiFi := new(WiFiHandle)
	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, ErrExpected)

	service := taskWifiPack.New(mockWiFi)
	names, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "get interfaces")

	mockWiFi.AssertExpectations(t)
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockWiFi := new(WiFiHandle)
	service := taskWifiPack.New(mockWiFi)

	assert.NotNil(t, service)
	assert.Equal(t, mockWiFi, service.WiFi)
}
