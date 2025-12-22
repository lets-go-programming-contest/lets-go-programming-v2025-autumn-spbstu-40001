package network

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
)

type WiFiManager interface {
	Interfaces() ([]*wifi.Interface, error)
	StationInfo(ifi *wifi.Interface) (*wifi.StationInfo, error)
}

type NetworkService struct {
	WiFi WiFiManager
}

func New(wifi WiFiManager) NetworkService {
	return NetworkService{WiFi: wifi}
}

func (service NetworkService) GetConnectedDevices() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("get interfaces: %w", err)
	}

	devices := make([]net.HardwareAddr, 0, len(interfaces))
	for _, iface := range interfaces {
		info, err := service.WiFi.StationInfo(iface)
		if err != nil {
			continue
		}

		if info.Connected {
			devices = append(devices, iface.HardwareAddr)
		}
	}

	return devices, nil
}

func (service NetworkService) GetInterfaceSpeeds() (map[string]int, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("get interfaces: %w", err)
	}

	speeds := make(map[string]int)
	for _, iface := range interfaces {
		info, err := service.WiFi.StationInfo(iface)
		if err != nil {
			continue
		}
		speeds[iface.Name] = info.TXBitrate / 1000000
	}

	return speeds, nil
}
