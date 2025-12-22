package wifi

import (
	"fmt"
	"net"

	mdwifi "github.com/mdlayher/wifi"
)

type WiFiHandle interface {
	Interfaces() ([]*mdwifi.Interface, error)
}

type Service struct {
	handle WiFiHandle
}

func New(handle WiFiHandle) *Service {
	return &Service{handle: handle}
}

func (s *Service) GetAddresses() ([]net.HardwareAddr, error) {
	ifaces, err := s.handle.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %w", err) // Добавить fmt.Errorf
	}

	addrs := make([]net.HardwareAddr, 0, len(ifaces))
	for _, iface := range ifaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (s *Service) GetNames() ([]string, error) {
	ifaces, err := s.handle.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get interfaces: %w", err) // Добавить fmt.Errorf
	}

	names := make([]string, 0, len(ifaces))
	for _, iface := range ifaces {
		names = append(names, iface.Name)
	}

	return names, nil
}
