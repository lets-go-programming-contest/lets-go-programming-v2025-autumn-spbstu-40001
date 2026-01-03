package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/PigoDog/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

const (
	selectAllNames = "SELECT name FROM users"
	selectDistinct = "SELECT DISTINCT name FROM users"
)

var errTest = errors.New("test error")

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	require.Equal(t, mockDB, service.DB)
}

func TestGetNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alex").
		AddRow("Maria")

	mock.ExpectQuery(selectAllNames).WillReturnRows(rows)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alex", "Maria"}, names)
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	mock.ExpectQuery(selectAllNames).WillReturnError(errTest)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "db query")
	require.Nil(t, names)
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery(selectAllNames).WillReturnRows(rows)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "rows scanning")
	require.Nil(t, names)
}

func TestGetNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alex")
	rows.RowError(0, errTest)
	mock.ExpectQuery(selectAllNames).WillReturnRows(rows)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "rows error")
	require.Nil(t, names)
}

func TestGetUniqueNames_Success(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alex").
		AddRow("Maria")

	mock.ExpectQuery(selectDistinct).WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alex", "Maria"}, names)
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	mock.ExpectQuery(selectDistinct).WillReturnError(errTest)

	names, err := service.GetUniqueNames()
	require.ErrorContains(t, err, "db query")
	require.Nil(t, names)
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery(selectDistinct).WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.ErrorContains(t, err, "rows scanning")
	require.Nil(t, names)
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	t.Cleanup(func() { _ = mockDB.Close() })

	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alex")
	rows.RowError(0, errTest)
	mock.ExpectQuery(selectDistinct).WillReturnRows(rows)

	names, err := service.GetUniqueNames()
	require.ErrorContains(t, err, "rows error")
	require.Nil(t, names)
}
