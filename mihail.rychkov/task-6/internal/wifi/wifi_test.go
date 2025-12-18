package wifi_test;

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

import "testing";
import "errors";
import "net";
import "github.com/mdlayher/wifi";
import "github.com/stretchr/testify/require";
import service "github.com/Rychmick/task-6/internal/wifi";

var errDefault = errors.New("something went wrong...");

var testCases = []struct {
	useGetNames    bool;
	names          []string;
	macs           []string;
	errExpectedMsg string;
	errExpected    error;
	errQuery       error;
	} {
	{false, []string{"device1"}, []string{"00:01:02:03:04:05"}, "", nil, nil},
	{false, nil, nil, "getting interfaces", errDefault, errDefault},
	{true, []string{"device1"}, []string{"00:01:02:03:04:05"}, "", nil, nil},
	{true, nil, nil, "getting interfaces", errDefault, errDefault},
};

func TestWiFi(t *testing.T) {
	t.Parallel();

	for _, testData := range testCases {
		mock := NewWiFiHandle(t);
		serviceObj := service.New(mock);
		expectError := (testData.errExpected != nil) || (testData.errExpectedMsg != "");

		require.Equal(t, len(testData.names), len(testData.macs));
		interfaces := make([]*wifi.Interface, 0, len(testData.macs))
		expectedMacs := make([]net.HardwareAddr, 0, );
		for i := range len(testData.macs) {
			parsedMAC, err := net.ParseMAC(testData.macs[i]);
			require.NoError(t, err);

			expectedMacs = append(expectedMacs, parsedMAC);
			interfaces = append(interfaces, &wifi.Interface{
				Index:        i,
				Name:         testData.names[i],
				HardwareAddr: parsedMAC,
			});
		}
		mock.On("Interfaces").Return(interfaces, testData.errQuery);
		var err error;
		if (testData.useGetNames) {
			names, err1 := serviceObj.GetNames();
			if (!expectError) {
				require.NoError(t, err);
				require.Equal(t, testData.names, names);
				continue;
			}
			err = err1;
		} else {
			macs, err1 := serviceObj.GetAddresses();
			if (!expectError) {
				require.NoError(t, err);
				require.Equal(t, expectedMacs, macs);
				continue;
			}
			err = err1;
		}
		if (testData.errExpected != nil) {
			require.ErrorIs(t, err, testData.errExpected);
		} else {
			require.Error(t, err);
		}
		require.ErrorContains(t, err, testData.errExpectedMsg);
	}
}
