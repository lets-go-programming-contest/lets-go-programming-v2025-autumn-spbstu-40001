package wifi_test

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

import (
	"errors"
	"fmt"
	"net"
	"testing"

	service "github.com/Rychmick/task-6/internal/wifi"
	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

func generateInterfaces() []*wifi.Interface {
	return []*wifi.Interface{getInterface("device1", "00:01:02:03:04:05")}
}

func getInterface(name, mac string) *wifi.Interface {
	parsedMAC, err := net.ParseMAC(mac)
	if err != nil {
		panic(err)
	}

	return &wifi.Interface{
		Name:         name,
		HardwareAddr: parsedMAC,
	}
}

func queryNames(t *testing.T, serviceObj service.WiFiService, ifaces []*wifi.Interface, checkResult bool) error {
	t.Helper()

	receivedNames, err := serviceObj.GetNames()

	if checkResult {
		require.NoError(t, err)

		for i, device := range ifaces {
			require.Equal(t, device.Name, receivedNames[i])
		}
	}

	return fmt.Errorf("received error: %w", err)
}

func queryAddresses(t *testing.T, serviceObj service.WiFiService, ifaces []*wifi.Interface, checkResult bool) error {
	t.Helper()

	recievedMacs, err := serviceObj.GetAddresses()

	if checkResult {
		require.NoError(t, err)

		for i, device := range ifaces {
			require.Equal(t, device.HardwareAddr, recievedMacs[i])
		}
	}

	return fmt.Errorf("received error: %w", err)
}

var errDefault = errors.New("something went wrong")

var testCases = []struct { //nolint:gochecknoglobals
	method         func(t *testing.T, serviceObj service.WiFiService, ifaces []*wifi.Interface, checkResult bool) error
	interfaces     []*wifi.Interface
	errExpectedMsg string
	errExpected    error
	errQuery       error
}{
	{queryAddresses, generateInterfaces(), "", nil, nil},
	{queryAddresses, nil, "getting interfaces", errDefault, errDefault},
	{queryNames, generateInterfaces(), "", nil, nil},
	{queryNames, nil, "getting interfaces", errDefault, errDefault},
}

func TestWiFi(t *testing.T) {
	t.Parallel()

	for i, testData := range testCases {
		t.Run(fmt.Sprintf("testcase #%d", i), func(t *testing.T) {
			t.Parallel()
			mock := NewWiFiHandle(t)
			serviceObj := service.New(mock)
			expectError := (testData.errExpected != nil) || (testData.errExpectedMsg != "")

			mock.On("Interfaces").Return(testData.interfaces, testData.errQuery)

			err := testData.method(t, serviceObj, testData.interfaces, !expectError)

			if !expectError {
				return
			}

			if testData.errExpected != nil {
				require.ErrorIs(t, err, testData.errExpected)
			}

			require.ErrorContains(t, err, testData.errExpectedMsg)
		})
	}
}
