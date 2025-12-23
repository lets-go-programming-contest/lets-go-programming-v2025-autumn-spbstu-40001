package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"

	wifiservice "github.com/Danil3352/task-6/internal/wifi"
)

func TestGetAddresses(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockHandle := new(WiFiHandleMock)

		addr1, _ := net.ParseMAC("00:11:22:33:44:55")
		addr2, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")

		mockInterfaces := []*wifi.Interface{
			{Index: 1, Name: "wlan0", HardwareAddr: addr1},
			{Index: 2, Name: "wlan1", HardwareAddr: addr2},
		}

		mockHandle.On("Interfaces").Return(mockInterfaces, nil)

		service := wifiservice.New(mockHandle)
		addrs, err := service.GetAddresses()

		assert.NoError(t, err)
		assert.Len(t, addrs, 2)
		assert.Equal(t, addr1, addrs[0])
		assert.Equal(t, addr2, addrs[1])

		mockHandle.AssertExpectations(t)
	})

	t.Run("error_from_handle", func(t *testing.T) {
		mockHandle := new(WiFiHandleMock)
		mockHandle.On("Interfaces").Return(nil, errors.New("low level error"))

		service := wifiservice.New(mockHandle)
		addrs, err := service.GetAddresses()

		assert.Error(t, err)
		assert.Nil(t, addrs)
		assert.Contains(t, err.Error(), "getting interfaces")

		mockHandle.AssertExpectations(t)
	})
}

func TestGetNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockHandle := new(WiFiHandleMock)

		mockInterfaces := []*wifi.Interface{
			{Name: "wlan0"},
			{Name: "eth0"},
		}

		mockHandle.On("Interfaces").Return(mockInterfaces, nil)

		service := wifiservice.New(mockHandle)
		names, err := service.GetNames()

		assert.NoError(t, err)
		assert.Equal(t, []string{"wlan0", "eth0"}, names)

		mockHandle.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockHandle := new(WiFiHandleMock)
		mockHandle.On("Interfaces").Return(nil, errors.New("failed to fetch"))

		service := wifiservice.New(mockHandle)
		names, err := service.GetNames()

		assert.Error(t, err)
		assert.Nil(t, names)

		mockHandle.AssertExpectations(t)
	})
}

func TestNew(t *testing.T) {
	mockHandle := new(WiFiHandleMock)
	service := wifiservice.New(mockHandle)

	assert.NotNil(t, service)
	assert.Equal(t, mockHandle, service.WiFi)
}
