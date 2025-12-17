package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	myWiFi "github.com/netwite/task-6/internal/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

var (
	ErrExpected   = errors.New("expected error")
	ErrNilPointer = errors.New("nil pointer")
)

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("success with mock", func(t *testing.T) {
		t.Parallel()

		mockWiFi := NewWiFiHandle(t)
		service := myWiFi.New(mockWiFi)

		assert.NotNil(t, service)
		assert.Equal(t, mockWiFi, service.WiFi)
	})

	t.Run("nil handler", func(t *testing.T) {
		t.Parallel()

		service := myWiFi.New(nil)
		assert.NotNil(t, service)
		assert.Nil(t, service.WiFi)
	})
}

func parseMAC(t *testing.T, s string) net.HardwareAddr {
	t.Helper()

	hwAddr, err := net.ParseMAC(s)
	require.NoError(t, err, "failed to parse MAC address: %s", s)
	return hwAddr
}

func TestGetAddresses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		mockReturnInterfaces []*wifi.Interface
		mockReturnError      error
		expectedAddrs        []net.HardwareAddr
		expectError          bool
		errorIs              error
	}{
		{
			name: "success with multiple interfaces",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "wlan0",
					HardwareAddr: parseMAC(t, "00:11:22:33:44:55"),
				},
				{
					Index:        2,
					Name:         "eth0",
					HardwareAddr: parseMAC(t, "aa:bb:cc:dd:ee:ff"),
				},
			},
			mockReturnError: nil,
			expectedAddrs: []net.HardwareAddr{
				parseMAC(t, "00:11:22:33:44:55"),
				parseMAC(t, "aa:bb:cc:dd:ee:ff"),
			},
			expectError: false,
		},
		{
			name:                 "error from interfaces method",
			mockReturnInterfaces: nil,
			mockReturnError:      ErrExpected,
			expectedAddrs:        nil,
			expectError:          true,
			errorIs:              ErrExpected,
		},
		{
			name:                 "empty interfaces list",
			mockReturnInterfaces: []*wifi.Interface{},
			mockReturnError:      nil,
			expectedAddrs:        []net.HardwareAddr{},
			expectError:          false,
		},
		{
			name: "interface with nil hardware address",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "wlan0",
					HardwareAddr: nil,
				},
				{
					Index:        2,
					Name:         "eth0",
					HardwareAddr: parseMAC(t, "aa:bb:cc:dd:ee:ff"),
				},
			},
			mockReturnError: nil,
			expectedAddrs: []net.HardwareAddr{
				nil,
				parseMAC(t, "aa:bb:cc:dd:ee:ff"),
			},
			expectError: false,
		},
		{
			name: "single interface",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "wlan0",
					HardwareAddr: parseMAC(t, "00:11:22:33:44:55"),
				},
			},
			mockReturnError: nil,
			expectedAddrs: []net.HardwareAddr{
				parseMAC(t, "00:11:22:33:44:55"),
			},
			expectError: false,
		},
		{
			name: "interface with zero MAC address",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "wlan0",
					HardwareAddr: parseMAC(t, "00:00:00:00:00:00"),
				},
			},
			mockReturnError: nil,
			expectedAddrs: []net.HardwareAddr{
				parseMAC(t, "00:00:00:00:00:00"),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockWiFi := NewWiFiHandle(t)
			service := myWiFi.New(mockWiFi)

			mockWiFi.On("Interfaces").Return(tt.mockReturnInterfaces, tt.mockReturnError).Once()

			addrs, err := service.GetAddresses()

			if tt.expectError {
				require.Error(t, err)
				if tt.errorIs != nil {
					assert.ErrorIs(t, err, tt.errorIs)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedAddrs, addrs)

				// Additional validation for non-nil addresses
				for _, addr := range addrs {
					if addr != nil {
						assert.NotEmpty(t, addr.String())
					}
				}
			}

			mockWiFi.AssertExpectations(t)
		})
	}
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                 string
		mockReturnInterfaces []*wifi.Interface
		mockReturnError      error
		expectedNames        []string
		expectError          bool
		errorIs              error
	}{
		{
			name: "success with multiple interfaces",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "wlan0",
					HardwareAddr: parseMAC(t, "00:11:22:33:44:55"),
				},
				{
					Index:        2,
					Name:         "eth0",
					HardwareAddr: parseMAC(t, "aa:bb:cc:dd:ee:ff"),
				},
				{
					Index:        3,
					Name:         "wlan1",
					HardwareAddr: parseMAC(t, "11:22:33:44:55:66"),
				},
			},
			mockReturnError: nil,
			expectedNames:   []string{"wlan0", "eth0", "wlan1"},
			expectError:     false,
		},
		{
			name:                 "error from interfaces method",
			mockReturnInterfaces: nil,
			mockReturnError:      ErrExpected,
			expectedNames:        nil,
			expectError:          true,
			errorIs:              ErrExpected,
		},
		{
			name:                 "empty interfaces list",
			mockReturnInterfaces: []*wifi.Interface{},
			mockReturnError:      nil,
			expectedNames:        []string{},
			expectError:          false,
		},
		{
			name: "interface with empty name",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "",
					HardwareAddr: parseMAC(t, "00:11:22:33:44:55"),
				},
				{
					Index:        2,
					Name:         "eth0",
					HardwareAddr: parseMAC(t, "aa:bb:cc:dd:ee:ff"),
				},
			},
			mockReturnError: nil,
			expectedNames:   []string{"", "eth0"},
			expectError:     false,
		},
		{
			name: "single interface",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "wlan0",
					HardwareAddr: parseMAC(t, "00:11:22:33:44:55"),
				},
			},
			mockReturnError: nil,
			expectedNames:   []string{"wlan0"},
			expectError:     false,
		},
		{
			name: "interface names with special characters",
			mockReturnInterfaces: []*wifi.Interface{
				{
					Index:        1,
					Name:         "wlan-0",
					HardwareAddr: parseMAC(t, "00:11:22:33:44:55"),
				},
				{
					Index:        2,
					Name:         "eth_1",
					HardwareAddr: parseMAC(t, "aa:bb:cc:dd:ee:ff"),
				},
			},
			mockReturnError: nil,
			expectedNames:   []string{"wlan-0", "eth_1"},
			expectError:     false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockWiFi := NewWiFiHandle(t)
			service := myWiFi.New(mockWiFi)

			mockWiFi.On("Interfaces").Return(tt.mockReturnInterfaces, tt.mockReturnError).Once()

			names, err := service.GetNames()

			if tt.expectError {
				require.Error(t, err)
				if tt.errorIs != nil {
					assert.ErrorIs(t, err, tt.errorIs)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedNames, names)

				assert.Len(t, names, len(tt.expectedNames))
			}

			mockWiFi.AssertExpectations(t)
		})
	}
}
