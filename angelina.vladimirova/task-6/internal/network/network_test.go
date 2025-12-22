package network_test

import (
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
	"github.com/verticalochka/task-6/internal/network"
)

func parseMAC(addr string) []byte {
	mac, _ := net.ParseMAC(addr)
	return mac
}

func TestGetConnectedDevices(t *testing.T) {
	t.Parallel()

	mockWiFi := newMockWiFiManager(t)
	service := network.New(mockWiFi)

	interfaces := []*wifi.Interface{
		{
			Name:         "wlan0",
			HardwareAddr: parseMAC("11:22:33:44:55:66"),
		},
		{
			Name:         "wlan1",
			HardwareAddr: parseMAC("aa:bb:cc:dd:ee:ff"),
		},
	}

	stationInfo1 := &wifi.StationInfo{Connected: true, TXBitrate: 300000000}
	stationInfo2 := &wifi.StationInfo{Connected: false, TXBitrate: 0}

	mockWiFi.On("Interfaces").Return(interfaces, nil)
	mockWiFi.On("StationInfo", interfaces[0]).Return(stationInfo1, nil)
	mockWiFi.On("StationInfo", interfaces[1]).Return(stationInfo2, nil)

	devices, err := service.GetConnectedDevices()
	require.NoError(t, err)
	require.Len(t, devices, 1)
	require.Equal(t, parseMAC("11:22:33:44:55:66"), devices[0])

	mockWiFi.On("Interfaces").Return([]*wifi.Interface{}, nil)

	devices, err = service.GetConnectedDevices()
	require.NoError(t, err)
	require.Empty(t, devices)
}

func TestGetInterfaceSpeeds(t *testing.T) {
	t.Parallel()

	mockWiFi := newMockWiFiManager(t)
	service := network.New(mockWiFi)

	interfaces := []*wifi.Interface{
		{
			Name:         "wifi0",
			HardwareAddr: parseMAC("11:22:33:44:55:66"),
		},
		{
			Name:         "wifi1",
			HardwareAddr: parseMAC("aa:bb:cc:dd:ee:ff"),
		},
	}

	stationInfo1 := &wifi.StationInfo{Connected: true, TXBitrate: 433300000}
	stationInfo2 := &wifi.StationInfo{Connected: true, TXBitrate: 866700000}

	mockWiFi.On("Interfaces").Return(interfaces, nil)
	mockWiFi.On("StationInfo", interfaces[0]).Return(stationInfo1, nil)
	mockWiFi.On("StationInfo", interfaces[1]).Return(stationInfo2, nil)

	speeds, err := service.GetInterfaceSpeeds()
	require.NoError(t, err)
	require.Equal(t, map[string]int{
		"wifi0": 433,
		"wifi1": 866,
	}, speeds)
}
