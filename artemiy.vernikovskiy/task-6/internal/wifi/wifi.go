package wifi

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/mock"
)

type WiFiInterfaces interface {
	Interfaces() ([]*wifi.Interface, error)
}

type WiFiService struct {
	WiFi WiFiInterfaces
}

func New(wifi WiFiInterfaces) WiFiService {
	return WiFiService{WiFi: wifi}
}

func (service WiFiService) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("getting interfaces: %w", err)
	}

	addrs := make([]net.HardwareAddr, 0, len(interfaces))

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service WiFiService) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("getting interfaces: %w", err)
	}

	names := make([]string, 0, len(interfaces))

	for _, iface := range interfaces {
		names = append(names, iface.Name)
	}

	return names, nil
}

type WiFiHandle struct {
	mock.Mock
}

func (m *WiFiHandle) Interfaces() ([]*wifi.Interface, error) {
	args := m.Called()

	var err error
	if args.Error(1) != nil {
		err = fmt.Errorf("mock error: %w", args.Error(1))
	}

	if args.Get(0) == nil {
		return nil, err
	}

	if ifaces, ok := args.Get(0).([]*wifi.Interface); ok {
		return ifaces, err
	}

	return nil, err
}
