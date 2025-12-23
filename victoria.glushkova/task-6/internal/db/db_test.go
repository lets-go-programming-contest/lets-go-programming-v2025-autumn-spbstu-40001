package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/vikaglushkova/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type WiFiHandleMock struct {
	mock.Mock
}

func (m *WiFiHandleMock) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()
	ifaces, _ := args.Get(0).([]*wifi.Interface)
	return ifaces, args.Error(1)
}

func NewWiFiHandle(t *testing.T) *WiFiHandleMock {
	t.Helper()
	return &WiFiHandleMock{}
}

func mustMAC(s string) net.HardwareAddr {
	m, err := net.ParseMAC(s)
	if err != nil {
		panic(err)
	}
	return m
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	ifaces := []*wifi.Interface{
		{
			Name:         "eth0",
			HardwareAddr: mustMAC("00:11:22:33:44:55"),
		},
		{
			Name:         "wlan0",
			HardwareAddr: mustMAC("aa:bb:cc:dd:ee:ff"),
		},
	}

	mockWiFi.On("Interfaces").Return(ifaces, nil)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{
		mustMAC("00:11:22:33:44:55"),
		mustMAC("aa:bb:cc:dd:ee:ff"),
	}, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestGetAddresses_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	mockWiFi.On("Interfaces").Return(nil, errors.New("iface error"))

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()

	require.Error(t, err)
	require.Nil(t, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	ifaces := []*wifi.Interface{
		{Name: "eth0"},
		{Name: "wlan0"},
	}

	mockWiFi.On("Interfaces").Return(ifaces, nil)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"eth0", "wlan0"}, names)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	mockWiFi.On("Interfaces").Return(nil, errors.New("iface error"))

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestGetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	require.Empty(t, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Empty(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestGetAddresses_NilMAC(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	ifaces := []*wifi.Interface{
		{
			Name:         "eth0",
			HardwareAddr: mustMAC("00:11:22:33:44:55"),
		},
		{
			Name:         "wlan0",
			HardwareAddr: nil,
		},
	}

	mockWiFi.On("Interfaces").Return(ifaces, nil)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	require.Len(t, addrs, 2)
	require.Equal(t, mustMAC("00:11:22:33:44:55"), addrs[0])
	require.Nil(t, addrs[1])

	mockWiFi.AssertExpectations(t)
}

func TestNewService(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)
	require.NotNil(t, service)

	mockWiFi.AssertNotCalled(t, "Interfaces")
}
