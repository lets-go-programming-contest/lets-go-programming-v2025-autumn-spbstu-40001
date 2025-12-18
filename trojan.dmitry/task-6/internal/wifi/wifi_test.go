package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/DimasFantomasA/task-6/internal/wifi"

	mdwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var errWiFiError = errors.New("wifi error")

func TestNew(t *testing.T) {
	t.Parallel() // Добавлено

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	require.NotNil(t, service)
	require.NotNil(t, service.WiFi)
}

func TestWiFiService_GetAddresses_Success(t *testing.T) {
	t.Parallel() // Добавлено

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*mdwifi.Interface{
		{HardwareAddr: mustMAC("00:11:22:33:44:55")},
		{HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff")},
	}, nil)

	result, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{
		mustMAC("00:11:22:33:44:55"),
		mustMAC("aa:bb:cc:dd:ee:ff"),
	}, result)
}

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	t.Parallel() // Добавлено

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	expectedErr := errWiFiError
	mockWiFi.On("Interfaces").Return(([]*mdwifi.Interface)(nil), expectedErr)

	result, err := service.GetAddresses()

	require.Error(t, err)
	require.Contains(t, err.Error(), "getting interfaces")
	require.Contains(t, err.Error(), expectedErr.Error())
	require.Nil(t, result)
}

func TestWiFiService_GetAddresses_Empty(t *testing.T) {
	t.Parallel() // Добавлено

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*mdwifi.Interface{}, nil)

	result, err := service.GetAddresses()

	require.NoError(t, err)
	require.Empty(t, result)
}

func TestWiFiService_GetNames_Success(t *testing.T) {
	t.Parallel() // Добавлено

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*mdwifi.Interface{
		{Name: "wlan0"},
		{Name: "eth0"},
		{Name: "wlan1"},
	}, nil)

	result, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "eth0", "wlan1"}, result)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	t.Parallel() // Добавлено

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	expectedErr := errWiFiError
	mockWiFi.On("Interfaces").Return(([]*mdwifi.Interface)(nil), expectedErr)

	result, err := service.GetNames()

	require.Error(t, err)
	require.Contains(t, err.Error(), "getting interfaces")
	require.Contains(t, err.Error(), expectedErr.Error())
	require.Nil(t, result)
}

func TestWiFiService_GetNames_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	mockWiFi.On("Interfaces").Return([]*mdwifi.Interface{}, nil)

	result, err := service.GetNames()

	require.NoError(t, err)
	require.Empty(t, result)
}

func mustMAC(s string) net.HardwareAddr {
	mac, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}

	return mac
}
