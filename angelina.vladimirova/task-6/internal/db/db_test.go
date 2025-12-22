package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/verticalochka/task-6/internal/db"
)

func TestGetNames(t *testing.T) {
	t.Parallel()

	dbMock, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbMock.Close()

	service := db.New(dbMock)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("alex").
		AddRow("maria")

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetNames()
	require.NoError(t, err)
	require.Equal(t, []string{"alex", "maria"}, result)
}

func TestGetNames_FailedQuery(t *testing.T) {
	t.Parallel()

	dbMock, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbMock.Close()

	service := db.New(dbMock)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errors.New("test error"))

	result, err := service.GetNames()
	require.ErrorContains(t, err, "db query")
	require.Nil(t, result)
}

func TestGetNames_BadScan(t *testing.T) {
	t.Parallel()

	dbMock, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbMock.Close()

	service := db.New(dbMock)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetNames()
	require.ErrorContains(t, err, "rows scanning")
	require.Nil(t, result)
}
