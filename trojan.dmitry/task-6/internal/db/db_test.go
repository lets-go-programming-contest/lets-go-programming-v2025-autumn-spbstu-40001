package db_test

import (
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"github.com/DimasFantomasA/task-6/internal/db"
)

// Тест для конструктора New
func TestNew(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	service := db.New(mockDB)

	require.NotNil(t, service)
	require.NotNil(t, service.DB)
}

func TestGetNames_Success(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow("Ivan").
			AddRow("Petr"))

	names, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Petr"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnError(errors.New("query error"))

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow(driver.Value(123)),
		)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errors.New("rows error"))

	mock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_Success(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow("Ivan").
			AddRow("Petr").
			AddRow("Ivan"))

	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	require.Equal(t, []string{"Ivan", "Petr"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_EmptyResult(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(sqlmock.NewRows([]string{"name"}))

	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	require.Empty(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(errors.New("query error"))

	names, err := service.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow(driver.Value(123)),
		)

	names, err := service.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsError(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	service := db.New(mockDB)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errors.New("rows error"))

	mock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	names, err := service.GetUniqueNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.NoError(t, mock.ExpectationsWereMet())
}
