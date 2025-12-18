package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Rychmick/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

const (
	queryUsual    = "SELECT name FROM users"
	queryDistinct = "SELECT DISTINCT name FROM users"
)

var (
	errDefault = errors.New("something went wrong")
	headings   = []string{"name"} //nolint:gochecknoglobals
)

var testCases = []struct { //nolint:gochecknoglobals
	useGetUnique   bool
	rows           *sqlmock.Rows
	names          []string
	errExpectedMsg string
	errExpected    error
	errQuery       error
}{
	{false, sqlmock.NewRows(headings).AddRow("1"), []string{"1"}, "", nil, nil},
	{false, sqlmock.NewRows(headings).AddRow("1"), nil, "db query", errDefault, errDefault},
	{false, sqlmock.NewRows(headings).AddRow(nil), nil, "rows scanning", nil, nil},
	{false, sqlmock.NewRows(headings).AddRow("1").RowError(0, errDefault), nil, "rows error", errDefault, nil},
	{true, sqlmock.NewRows(headings).AddRow("1"), []string{"1"}, "", nil, nil},
	{true, sqlmock.NewRows(headings).AddRow("1"), nil, "db query", errDefault, errDefault},
	{true, sqlmock.NewRows(headings).AddRow(nil), nil, "rows scanning", nil, nil},
	{true, sqlmock.NewRows(headings).AddRow("1").RowError(0, errDefault), nil, "rows error", errDefault, nil},
}

func TestDatabase(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		require.NoError(t, err)
	}
	defer mockDB.Close()

	libDB := db.New(mockDB)

	for _, testData := range testCases {
		var names []string

		if testData.useGetUnique {
			mock.ExpectQuery(queryDistinct).WillReturnRows(testData.rows).WillReturnError(testData.errQuery)

			names, err = libDB.GetUniqueNames()
		} else {
			mock.ExpectQuery(queryUsual).WillReturnRows(testData.rows).WillReturnError(testData.errQuery)

			names, err = libDB.GetNames()
		}

		require.NoError(t, mock.ExpectationsWereMet())

		if (testData.errExpected != nil) || (testData.errExpectedMsg != "") {
			if testData.errExpected != nil {
				require.ErrorIs(t, err, testData.errExpected)
			} else {
				require.Error(t, err)
			}

			require.ErrorContains(t, err, testData.errExpectedMsg)
			require.Empty(t, names)

			continue
		}

		require.NoError(t, err)
		require.Equal(t, testData.names, names)
	}
}
