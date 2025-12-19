package db_test

import (
	"errors"
	"fmt"
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

func generateRows() *sqlmock.Rows {
	return sqlmock.NewRows(headings).AddRow("1")
}

func generateErrRows() *sqlmock.Rows {
	return generateRows().RowError(0, errDefault)
}

func generateNilRows() *sqlmock.Rows {
	return generateRows().AddRow(nil)
}

func usualQuery(mock sqlmock.Sqlmock, service db.DBService, rows *sqlmock.Rows, errQuery error) ([]string, error) {
	mock.ExpectQuery(queryUsual).WillReturnRows(rows).WillReturnError(errQuery)

	names, err := service.GetNames()
	if err != nil {
		return names, fmt.Errorf("received error: %w", err)
	}

	return names, nil
}

func distinctQuery(mock sqlmock.Sqlmock, service db.DBService, rows *sqlmock.Rows, errQuery error) ([]string, error) {
	mock.ExpectQuery(queryDistinct).WillReturnRows(rows).WillReturnError(errQuery)

	names, err := service.GetUniqueNames()
	if err != nil {
		return names, fmt.Errorf("received error: %w", err)
	}

	return names, nil
}

var testCases = []struct { //nolint:gochecknoglobals
	method         func(mock sqlmock.Sqlmock, service db.DBService, rows *sqlmock.Rows, errQuery error) ([]string, error)
	rows           *sqlmock.Rows
	errQuery       error
	names          []string
	errExpectedMsg string
	errExpected    error
}{
	{usualQuery, generateRows(), nil, []string{"1"}, "", nil},
	{usualQuery, generateRows(), errDefault, nil, "db query", errDefault},
	{usualQuery, generateErrRows(), nil, nil, "rows error", errDefault},
	{usualQuery, generateNilRows(), nil, nil, "rows scanning", nil},
	{distinctQuery, generateRows(), nil, []string{"1"}, "", nil},
	{distinctQuery, generateRows(), errDefault, nil, "db query", errDefault},
	{distinctQuery, generateErrRows(), nil, nil, "rows error", errDefault},
	{distinctQuery, generateNilRows(), nil, nil, "rows scanning", nil},
}

func TestDatabase(t *testing.T) {
	t.Parallel()

	for i, testData := range testCases {
		t.Run(fmt.Sprintf("testcase #%d", i), func(t *testing.T) {
			t.Parallel()

			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			defer mockDB.Close()

			service := db.New(mockDB)

			names, err := testData.method(mock, service, testData.rows, testData.errQuery)

			require.NoError(t, mock.ExpectationsWereMet())

			if (testData.errExpected != nil) || (testData.errExpectedMsg != "") {
				if testData.errExpected != nil {
					require.ErrorIs(t, err, testData.errExpected)
				}

				require.ErrorContains(t, err, testData.errExpectedMsg)
				require.Empty(t, names)

				return
			}

			require.NoError(t, err)
			require.Equal(t, testData.names, names)
		})
	}
}
