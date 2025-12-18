package db_test

import (
	"errors"
	"testing"

	"github.com/DimasFantomasA/task-6/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestDBService_GetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	service := db.New(mockDB)

	tests := []struct {
		name        string
		dbRows      []string
		dbErr       error
		expected    []string
		expectError bool
	}{
		{
			name:     "success",
			dbRows:   []string{"Ivan", "Petr"},
			expected: []string{"Ivan", "Petr"},
		},
		{
			name:        "db error",
			dbErr:       errors.New("db error"),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("SELECT name FROM users").
				WillReturnRows(mockRows(tt.dbRows)).
				WillReturnError(tt.dbErr)

			result, err := service.GetNames()

			if tt.expectError {
				require.Error(t, err)
				require.Nil(t, result)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestDBService_GetUniqueNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(mockRows([]string{"Ivan", "Petr"}))

	result, err := service.GetUniqueNames()

	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Petr"}, result)
}

func mockRows(values []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, v := range values {
		rows.AddRow(v)
	}
	return rows
}
