package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/verticalochka/task-6/internal/db"
)

const selectAllNames = "SELECT name FROM users"

var ErrTest = errors.New("test error")

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.New(mockDB)
	require.Equal(t, mockDB, service.DB)
}

func TestGetNames_Successful(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alex").
		AddRow("Maria")

	mock.ExpectQuery(selectAllNames).WillReturnRows(rows)

	names, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"Alex", "Maria"}, names)
}

func TestGetNames_FailedQuery(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.DBService{DB: mockDB}

	mock.ExpectQuery(selectAllNames).WillReturnError(ErrTest)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "db query")
	require.Nil(t, names)
}

func TestGetNames_BadScan(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery(selectAllNames).WillReturnRows(rows)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "rows scanning")
	require.Nil(t, names)
}

func TestGetNames_ProblemRows(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	service := db.DBService{DB: mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Peter")
	rows.RowError(0, ErrTest)
	mock.ExpectQuery(selectAllNames).WillReturnRows(rows)

	names, err := service.GetNames()
	require.ErrorContains(t, err, "rows error")
	require.Nil(t, names)
}
