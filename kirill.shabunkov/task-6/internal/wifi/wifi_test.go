package wifi_test

import (
	"errors"
	"net"
	"testing"

	"kirill.shabunkov/task-6/internal/wifi"

	mdwifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var errSystem = errors.New("system error")

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(MockWiFiHandle)
		mac, _ := net.ParseMAC("00:00:5e:00:53:01")
		interfaces := []*mdwifi.Interface{{HardwareAddr: mac}}

		mock.On("Interfaces").Return(interfaces, nil)

		service := wifi.New(mock)
		addrs, err := service.GetAddresses()

		require.NoError(t, err)
		assert.Len(t, addrs, 1)
		assert.Equal(t, mac, addrs[0])
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		mock := new(MockWiFiHandle)
		mock.On("Interfaces").Return(nil, errSystem)

		service := wifi.New(mock)
		_, err := service.GetAddresses()

		require.Error(t, err)
	})
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(MockWiFiHandle)
		interfaces := []*mdwifi.Interface{{Name: "wlan0"}}

		mock.On("Interfaces").Return(interfaces, nil)

		service := wifi.New(mock)
		names, err := service.GetNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"wlan0"}, names)
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		mock := new(MockWiFiHandle)
		mock.On("Interfaces").Return(nil, errSystem)

		service := wifi.New(mock)
		_, err := service.GetNames()

		require.Error(t, err)
	})
}
