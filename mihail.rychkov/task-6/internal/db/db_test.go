package db_test;

import "testing";
import "errors";
import "github.com/DATA-DOG/go-sqlmock";
import "github.com/stretchr/testify/require";
import "github.com/Rychmick/task-6/internal/db";

const queryUsual = "SELECT name FROM users";
const queryDistinct = "SELECT DISTINCT name FROM users";

var defaultError = errors.New("something went wrong...");
var headings = []string{"name"};

var getNamesCases = []struct {
	useGetUnique bool;
	rows           *sqlmock.Rows;
	names          []string;
	errExpectedMsg string;
	errExpected    error;
	errQuery       error;
	} {
	{false, sqlmock.NewRows(headings).AddRow("1"), []string{"1"}, "", nil, nil},
	{false, sqlmock.NewRows(headings).AddRow("1"), nil, "db query", defaultError, defaultError},
	{false, sqlmock.NewRows(headings).AddRow(nil), nil, "rows scanning", nil, nil},
	{false, sqlmock.NewRows(headings).AddRow("1").RowError(0, defaultError), nil, "rows error", defaultError, nil},
	{true, sqlmock.NewRows(headings).AddRow("1"), []string{"1"}, "", nil, nil},
	{true, sqlmock.NewRows(headings).AddRow("1"), nil, "db query", defaultError, defaultError},
	{true, sqlmock.NewRows(headings).AddRow(nil), nil, "rows scanning", nil, nil},
	{true, sqlmock.NewRows(headings).AddRow("1").RowError(0, defaultError), nil, "rows error", defaultError, nil},
};

func TestGetNames(t *testing.T) {
	t.Parallel();

	mockDB, mock, err := sqlmock.New();
	if (err != nil) {
		require.NoError(t, err);
	}
	defer mockDB.Close();

	libDB := db.New(mockDB);
	for _, testData := range(getNamesCases) {
		var names []string;
		if (testData.useGetUnique) {
			mock.ExpectQuery(queryDistinct).WillReturnRows(testData.rows).WillReturnError(testData.errQuery);
			names, err = libDB.GetUniqueNames();
		} else {
			mock.ExpectQuery(queryUsual).WillReturnRows(testData.rows).WillReturnError(testData.errQuery);
			names, err = libDB.GetNames();
		}
		require.NoError(t, mock.ExpectationsWereMet());
		if ((testData.errExpected != nil) || (testData.errExpectedMsg != "")) {
			if (testData.errExpected != nil) {
				require.ErrorIs(t, err, testData.errExpected);
			} else {
				require.Error(t, err);
			}
			require.ErrorContains(t, err, testData.errExpectedMsg);
			require.Empty(t, names);
			continue;
		}
		require.NoError(t, err);
		require.Equal(t, testData.names, names);
	}
}
