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

var (
	ErrExpected      = errors.New("expected error")
	ErrGettingIF     = errors.New("getting interfaces")
	ErrGetInterfaces = errors.New("get interfaces")
)

func TestWiFiServiceGetAddressesSuccess(t *testing.T) {
	const numberOfData = 3

	t.Parallel()

	mockWiFi := new(WiFiHandle)

	hwAddrs := []net.HardwareAddr{
		mustParseMAC("00:11:22:33:44:55"),
		mustParseMAC("aa:bb:cc:dd:ee:ff"),
		mustParseMAC("aa:bb:cc:dd:ee:ff"),
	}

	interfaces := []*wifi.Interface{
		{Name: "wlan0", HardwareAddr: hwAddrs[0]},
		{Name: "wlan1", HardwareAddr: hwAddrs[1]},
		{Name: "wlan2", HardwareAddr: hwAddrs[2]},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	service := taskWifiPack.New(mockWiFi)
	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	assert.Len(t, addrs, numberOfData)

	for i, addr := range hwAddrs {
		assert.Equal(t, addr, addrs[i])
	}

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
	assert.Contains(t, err.Error(), ErrGettingIF.Error())

	mockWiFi.AssertExpectations(t)
}

func TestWiFiServiceGetNamesSuccess(t *testing.T) {
	const numberOfData = 3

	t.Parallel()

	mockWiFi := new(WiFiHandle)

	ifNames := []string{"wlan0", "wlan1", "eth0"}
	hwAddr := mustParseMAC("13:37:de:ad:be:ef")
	interfaces := []*wifi.Interface{
		{Name: ifNames[0], HardwareAddr: hwAddr},
		{Name: ifNames[1]},
		{Name: ifNames[2]},
	}

	mockWiFi.On("Interfaces").Return(interfaces, nil)

	service := taskWifiPack.New(mockWiFi)
	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Len(t, names, numberOfData)
	assert.Equal(t, ifNames, names)

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
	assert.Contains(t, err.Error(), ErrGetInterfaces.Error())

	mockWiFi.AssertExpectations(t)
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockWiFi := new(WiFiHandle)
	service := taskWifiPack.New(mockWiFi)

	assert.NotNil(t, service)
	assert.Equal(t, mockWiFi, service.WiFi)
}

// if you want arrays, then we also need this
// (or i could make for loop, but this is more universal (i pretend that there is no error to throw)).
func mustParseMAC(addr string) net.HardwareAddr {
	hw, _ := net.ParseMAC(addr)

	return hw
}
