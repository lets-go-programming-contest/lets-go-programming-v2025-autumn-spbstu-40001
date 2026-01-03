package wifi_test

import (
	"errors"
	"net"
	"testing"

	myWifi "github.com/PigoDog/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

var errTest = errors.New("test error")

type wifiMock struct {
	interfaces []*wifi.Interface
	err        error
}

func (m *wifiMock) Interfaces() ([]*wifi.Interface, error) {
	return m.interfaces, m.err
}

func newService(mock *wifiMock) myWifi.WiFiService {
	return myWifi.New(mock)
}

func iface(name string, addr net.HardwareAddr) *wifi.Interface {
	return &wifi.Interface{
		Name:         name,
		HardwareAddr: addr,
	}
}

func hwAddr(b ...byte) net.HardwareAddr {
	return net.HardwareAddr(b)
}

func TestNew(t *testing.T) {
	t.Parallel()

	mock := &wifiMock{}
	service := myWifi.New(mock)

	require.Equal(t, mock, service.WiFi)
}

func TestGetAddresses_Success(t *testing.T) {
	t.Parallel()

	mock := &wifiMock{
		interfaces: []*wifi.Interface{
			iface("wlan0", hwAddr(0x00, 0x11, 0x22)),
			iface("wlan1", hwAddr(0x33, 0x44, 0x55)),
		},
	}

	service := newService(mock)

	addrs, err := service.GetAddresses()
	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{
		hwAddr(0x00, 0x11, 0x22),
		hwAddr(0x33, 0x44, 0x55),
	}, addrs)
}

func TestGetAddresses_InterfaceError(t *testing.T) {
	t.Parallel()

	mock := &wifiMock{
		err: errTest,
	}

	service := newService(mock)

	addrs, err := service.GetAddresses()
	require.ErrorContains(t, err, "getting interfaces")
	require.Nil(t, addrs)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mock := &wifiMock{
		interfaces: []*wifi.Interface{
			iface("wlan0", nil),
			iface("wlan1", nil),
		},
	}

	service := newService(mock)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"wlan0", "wlan1"}, names)
}

func TestGetNames_InterfaceError(t *testing.T) {
	t.Parallel()

	mock := &wifiMock{
		err: errTest,
	}

	service := newService(mock)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "getting interfaces")
	require.Nil(t, names)
}
