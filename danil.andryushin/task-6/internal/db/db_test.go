package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/atroxxxxxx/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

var errExpected = errors.New("error expected")

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	service := db.New(mockDB)
	require.Equal(t, mockDB, service.DB)
}

func TestDBService_GetNames(test *testing.T) {
	test.Parallel()

	tests := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectError   bool
		errorContains string
		expected      []string
	}{
		{
			name: "success",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("AA").AddRow("aa")
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"AA", "aa"},
		},
		{
			name: "query error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT name FROM users").WillReturnError(errExpected)
			},
			expectError:   true,
			errorContains: "db query",
		},
		{
			name: "rows error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("AA").RowError(0, errExpected)
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows error",
		},
		{
			name: "scan error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows scanning",
		},
	}
	for _, t := range tests {
		test.Run(t.name, func(t2 *testing.T) {
			t2.Parallel()

			dbConn, mock, err := sqlmock.New()
			require.NoError(t2, err)

			defer dbConn.Close()

			service := db.New(dbConn)

			t.mockSetup(mock)

			result, err := service.GetNames()
			if t.expectError {
				require.Error(t2, err)
				require.Contains(t2, err.Error(), t.errorContains)
			} else {
				require.NoError(t2, err)
				require.Equal(t2, t.expected, result)
			}
		})
	}
}

func TestDBService_GetUniqueNames(test *testing.T) {
	test.Parallel()

	tests := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectError   bool
		errorContains string
		expected      []string
	}{
		{
			name: "success",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("AA").AddRow("Ba")
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"AA", "Ba"},
		},
		{
			name: "query error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errExpected)
			},
			expectError:   true,
			errorContains: "db query",
		},
		{
			name: "rows error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow("AA").RowError(0, errExpected)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows error",
		},
		{
			name: "scan error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows scanning",
		},
	}

	for _, t := range tests {
		test.Run(t.name, func(t2 *testing.T) {
			t2.Parallel()

			dbConn, mock, err := sqlmock.New()
			require.NoError(t2, err)

			defer dbConn.Close()

			service := db.New(dbConn)

			t.mockSetup(mock)

			result, err := service.GetUniqueNames()
			if t.expectError {
				require.Error(t2, err)
				require.Contains(t2, err.Error(), t.errorContains)
			} else {
				require.NoError(t2, err)
				require.Equal(t2, t.expected, result)
			}

			require.NoError(t2, mock.ExpectationsWereMet())
		})
	}
}
