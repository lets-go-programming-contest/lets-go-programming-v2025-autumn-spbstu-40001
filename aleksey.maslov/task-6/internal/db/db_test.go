package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"github.com/A1exMas1ov/task-6/internal/db"
)

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	require.Equal(t, mockDB, service.DB)
}

var errExpected = errors.New("expected error")

func TestGetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectError   bool
		errorContains string
		expected      []string
	}{
		{
			name: "success",
			mockSetup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Alex").
					AddRow("Kate")
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"Alex", "Kate"},
		},
		{
			name: "query error",
			mockSetup: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT name FROM users").WillReturnError(errExpected)
			},
			expectError:   true,
			errorContains: "db query",
		},
		{
			name: "scan error",
			mockSetup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows scanning",
		},
		{
			name: "rows error",
			mockSetup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Alex").
					RowError(0, errExpected)
				m.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dbConn, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer dbConn.Close()

			service := db.New(dbConn)

			tt.mockSetup(mock)

			result, err := service.GetNames()

			if tt.expectError {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorContains)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectError   bool
		errorContains string
		expected      []string
	}{
		{
			name: "success",
			mockSetup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Alex").
					AddRow("Kate")
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expected: []string{"Alex", "Kate"},
		},
		{
			name: "query error",
			mockSetup: func(m sqlmock.Sqlmock) {
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errExpected)
			},
			expectError:   true,
			errorContains: "db query",
		},
		{
			name: "scan error",
			mockSetup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows scanning",
		},
		{
			name: "rows error",
			mockSetup: func(m sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"name"}).
					AddRow("Alex").
					RowError(0, errExpected)
				m.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)
			},
			expectError:   true,
			errorContains: "rows error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			dbConn, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer dbConn.Close()

			service := db.New(dbConn)

			tt.mockSetup(mock)

			result, err := service.GetUniqueNames()

			if tt.expectError {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errorContains)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, result)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
