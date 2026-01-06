package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	myWifi "github.com/A1exCRE/task-6/internal/wifi"
)

var errTest = errors.New("test error")

type testScenario struct {
	macAddresses []string
	expectedErr  error
}

func TestWiFiNew(t *testing.T) {
	t.Parallel()

	mockHandle := NewWiFiHandle(t)
	service := myWifi.New(mockHandle)
	require.Equal(t, mockHandle, service.WiFi)
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	scenarios := []testScenario{
		{macAddresses: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"}},
		{macAddresses: []string{}},
		{expectedErr: errTest},
	}

	mockHandle := NewWiFiHandle(t)
	service := myWifi.WiFiService{WiFi: mockHandle}

	for idx, sc := range scenarios {
		mockHandle.On("Interfaces").Unset()
		mockHandle.On("Interfaces").Return(createTestInterfaces(sc.macAddresses), sc.expectedErr)

		resultAddrs, err := service.GetAddresses()

		if sc.expectedErr != nil {
			require.ErrorIs(t, err, sc.expectedErr, "scenario %d", idx)
			require.ErrorContains(t, err, "getting interfaces", "scenario %d", idx)
			require.Nil(t, resultAddrs, "scenario %d", idx)

			continue
		}

		require.NoError(t, err, "scenario %d", idx)
		require.Equal(t, parseMACAddresses(sc.macAddresses), resultAddrs, "scenario %d", idx)
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	scenarios := []testScenario{
		{macAddresses: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"}},
		{macAddresses: []string{}},
		{expectedErr: errTest},
	}

	mockHandle := NewWiFiHandle(t)
	service := myWifi.WiFiService{WiFi: mockHandle}

	for idx, sc := range scenarios {
		mockHandle.On("Interfaces").Unset()
		mockHandle.On("Interfaces").Return(createTestInterfaces(sc.macAddresses), sc.expectedErr)

		resultNames, err := service.GetNames()

		if sc.expectedErr != nil {
			require.ErrorIs(t, err, sc.expectedErr, "scenario %d", idx)
			require.ErrorContains(t, err, "getting interfaces", "scenario %d", idx)
			require.Nil(t, resultNames, "scenario %d", idx)

			continue
		}

		require.NoError(t, err, "scenario %d", idx)
		require.Equal(t, generateExpectedNames(sc.macAddresses), resultNames, "scenario %d", idx)
	}
}

func generateExpectedNames(macAddrs []string) []string {
	names := make([]string, 0, len(macAddrs))
	for i := range macAddrs {
		names = append(names, fmt.Sprintf("wlan%d", i+1))
	}

	return names
}

func createTestInterfaces(macAddrs []string) []*wifi.Interface {
	ifaces := make([]*wifi.Interface, 0, len(macAddrs))

	for i, macStr := range macAddrs {
		hwAddr := parseMAC(macStr)
		if hwAddr == nil {
			continue
		}

		ifaces = append(ifaces, &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("wlan%d", i+1),
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		})
	}

	return ifaces
}

func parseMACAddresses(macStrs []string) []net.HardwareAddr {
	result := make([]net.HardwareAddr, 0, len(macStrs))
	for _, s := range macStrs {
		result = append(result, parseMAC(s))
	}

	return result
}

func parseMAC(macStr string) net.HardwareAddr {
	hw, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}

	return hw
}
