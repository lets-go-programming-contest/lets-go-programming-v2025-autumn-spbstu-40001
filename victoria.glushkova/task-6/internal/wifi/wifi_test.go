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

type WirelessInterfaceManagerMock struct {
	mock.Mock
}

func (m *WirelessInterfaceManagerMock) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()
	ifaceList, _ := args.Get(0).([]*wifi.Interface)
	return ifaceList, args.Error(1)
}

func createWirelessInterfaceManagerMock(t *testing.T) *WirelessInterfaceManagerMock {
	t.Helper()
	return &WirelessInterfaceManagerMock{}
}

func parseMACAddress(address string) net.HardwareAddr {
	mac, err := net.ParseMAC(address)
	if err != nil {
		panic(err)
	}
	return mac
}

func TestConstructorNew() {
	t := &testing.T{}
	mockManager := createWirelessInterfaceManagerMock(t)

	service := wifi.New(mockManager)
	require.NotNil(t, service)

	mockManager.AssertNotCalled(t, "Interfaces")
}

func TestWirelessServiceAddressRetrieval(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	interfaceList := []*wifi.Interface{
		{
			Name:         "wlan0",
			HardwareAddr: parseMACAddress("00:11:22:33:44:55"),
		},
		{
			Name:         "wlan1",
			HardwareAddr: parseMACAddress("aa:bb:cc:dd:ee:ff"),
		},
	}

	mockManager.On("Interfaces").Return(interfaceList, nil)

	wirelessService := wifi.New(mockManager)

	addresses, err := wirelessService.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{
		parseMACAddress("00:11:22:33:44:55"),
		parseMACAddress("aa:bb:cc:dd:ee:ff"),
	}, addresses)

	mockManager.AssertExpectations(t)
}

func TestWirelessServiceAddressRetrievalFailure(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	interfaceError := errors.New("system call failed")

	mockManager.On("Interfaces").Return(nil, interfaceError)

	wirelessService := wifi.New(mockManager)

	addresses, err := wirelessService.GetAddresses()

	require.Error(t, err)
	require.Contains(t, err.Error(), "getting interfaces")
	require.Contains(t, err.Error(), interfaceError.Error())
	require.Nil(t, addresses)

	mockManager.AssertExpectations(t)
}

func TestWirelessServiceNameRetrieval(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	interfaceList := []*wifi.Interface{
		{Name: "eth0"},
		{Name: "wlan0"},
		{Name: "wlan1"},
	}

	mockManager.On("Interfaces").Return(interfaceList, nil)

	wirelessService := wifi.New(mockManager)

	names, err := wirelessService.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"eth0", "wlan0", "wlan1"}, names)

	mockManager.AssertExpectations(t)
}

func TestWirelessServiceNameRetrievalFailure(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	accessError := errors.New("permission denied")

	mockManager.On("Interfaces").Return(nil, accessError)

	wirelessService := wifi.New(mockManager)

	names, err := wirelessService.GetNames()

	require.Error(t, err)
	require.Contains(t, err.Error(), "getting interfaces")
	require.Contains(t, err.Error(), accessError.Error())
	require.Nil(t, names)

	mockManager.AssertExpectations(t)
}

func TestWirelessServiceEmptyInterfaceList(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	emptyList := []*wifi.Interface{}

	mockManager.On("Interfaces").Return(emptyList, nil).Twice()

	wirelessService := wifi.New(mockManager)

	addresses, addrErr := wirelessService.GetAddresses()
	require.NoError(t, addrErr)
	require.Empty(t, addresses)

	names, nameErr := wirelessService.GetNames()
	require.NoError(t, nameErr)
	require.Empty(t, names)

	mockManager.AssertExpectations(t)
	mockManager.AssertNumberOfCalls(t, "Interfaces", 2)
}

func TestWirelessServiceInterfaceWithNilAddress(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	interfaceList := []*wifi.Interface{
		{
			Name:         "interface1",
			HardwareAddr: parseMACAddress("11:22:33:44:55:66"),
		},
		{
			Name:         "interface2",
			HardwareAddr: nil,
		},
	}

	mockManager.On("Interfaces").Return(interfaceList, nil)

	wirelessService := wifi.New(mockManager)

	addresses, err := wirelessService.GetAddresses()

	require.NoError(t, err)
	require.Len(t, addresses, 2)
	require.Equal(t, parseMACAddress("11:22:33:44:55:66"), addresses[0])
	require.Nil(t, addresses[1])

	mockManager.AssertExpectations(t)
}

func TestWirelessServiceConsecutiveMethodCalls(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	interfaceList := []*wifi.Interface{
		{
			Name:         "test0",
			HardwareAddr: parseMACAddress("01:23:45:67:89:ab"),
		},
	}

	mockManager.On("Interfaces").Return(interfaceList, nil).Twice()

	wirelessService := wifi.New(mockManager)

	addresses, addrErr := wirelessService.GetAddresses()
	require.NoError(t, addrErr)
	require.Len(t, addresses, 1)

	names, nameErr := wirelessService.GetNames()
	require.NoError(t, nameErr)
	require.Equal(t, []string{"test0"}, names)

	mockManager.AssertExpectations(t)
	mockManager.AssertNumberOfCalls(t, "Interfaces", 2)
}

func TestWirelessServiceMultipleInterfaces(t *testing.T) {
	t.Parallel()

	mockManager := createWirelessInterfaceManagerMock(t)
	interfaceList := []*wifi.Interface{
		{
			Name:         "wifi0",
			HardwareAddr: parseMACAddress("de:ad:be:ef:ca:fe"),
		},
		{
			Name:         "wifi1",
			HardwareAddr: parseMACAddress("fe:ed:fa:ce:ca:fe"),
		},
		{
			Name:         "wifi2",
			HardwareAddr: parseMACAddress("ca:fe:ba:be:de:ad"),
		},
	}

	mockManager.On("Interfaces").Return(interfaceList, nil)

	wirelessService := wifi.New(mockManager)

	addresses, addrErr := wirelessService.GetAddresses()
	require.NoError(t, addrErr)
	require.Len(t, addresses, 3)

	names, nameErr := wirelessService.GetNames()
	require.NoError(t, nameErr)
	require.Equal(t, []string{"wifi0", "wifi1", "wifi2"}, names)

	mockManager.AssertExpectations(t)
}

func TestWirelessServiceInterfaceImplementation() {
	var _ wifi.WiFiHandle = (*WirelessInterfaceManagerMock)(nil)

	service := wifi.WiFiService{}
	require.NotNil(t, service)
}
