package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	wifipkg "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/vikaglushkova/task-6/internal/wifi"
)

var errInterface = errors.New("interface error")

type WiFiHandleMock struct {
	mock.Mock
}

func (m *WiFiHandleMock) Interfaces() ([]*wifipkg.Interface, error) {
	args := m.Called()
	ifaces, _ := args.Get(0).([]*wifipkg.Interface)

	err := args.Error(1)
	if err != nil {
		return ifaces, fmt.Errorf("mock error: %w", err)
	}

	return ifaces, nil
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

func TestGetAddresses_Success(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	ifaces := []*wifipkg.Interface{
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
	mockWiFi.On("Interfaces").Return(nil, errInterface)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()

	require.Error(t, err)
	require.Contains(t, err.Error(), "getting interfaces")
	require.Nil(t, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestGetAddresses_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	mockWiFi.On("Interfaces").Return([]*wifipkg.Interface{}, nil)

	service := wifi.New(mockWiFi)

	addrs, err := service.GetAddresses()

	require.NoError(t, err)
	require.Empty(t, addrs)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	ifaces := []*wifipkg.Interface{
		{Name: "eth0"},
		{Name: "wlan0"},
		{Name: "wlan1"},
	}

	mockWiFi.On("Interfaces").Return(ifaces, nil)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"eth0", "wlan0", "wlan1"}, names)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames_Error(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	mockWiFi.On("Interfaces").Return(nil, errInterface)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Contains(t, err.Error(), "getting interfaces")
	require.Nil(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestGetNames_Empty(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	mockWiFi.On("Interfaces").Return([]*wifipkg.Interface{}, nil)

	service := wifi.New(mockWiFi)

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Empty(t, names)

	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_Constructor(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	service := wifi.New(mockWiFi)

	require.NotNil(t, service)
	mockWiFi.AssertNotCalled(t, "Interfaces")
}

func TestWiFiService_WithNilMACAddress(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	ifaces := []*wifipkg.Interface{
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

func TestWiFiService_MultipleCalls(t *testing.T) {
	t.Parallel()

	mockWiFi := NewWiFiHandle(t)
	ifaces := []*wifipkg.Interface{
		{
			Name:         "test0",
			HardwareAddr: mustMAC("01:23:45:67:89:ab"),
		},
	}

	mockWiFi.On("Interfaces").Return(ifaces, nil).Twice()

	service := wifi.New(mockWiFi)

	addrs, addrErr := service.GetAddresses()
	require.NoError(t, addrErr)
	require.Len(t, addrs, 1)

	names, nameErr := service.GetNames()
	require.NoError(t, nameErr)
	require.Equal(t, []string{"test0"}, names)

	mockWiFi.AssertExpectations(t)
	mockWiFi.AssertNumberOfCalls(t, "Interfaces", 2)
}
