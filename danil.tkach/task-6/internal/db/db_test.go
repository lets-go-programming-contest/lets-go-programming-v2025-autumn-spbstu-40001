package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	dbservice "github.com/Danil3352/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errInternalConn = errors.New("internal connection error")
	errScanFailed   = errors.New("scan failed")
	errFatalDB      = errors.New("fatal db error")
	errInterrupted  = errors.New("interrupted")
)

func TestGetNames_SuccessCase(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Yana").AddRow("Egor")
	mock.ExpectQuery("Select name from users").WillReturnRows(rows)

	names, executionErr := service.GetNames()

	require.NoError(t, executionErr)
	assert.Equal(t, []string{"Yana", "Egor"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_QueryError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)

	mock.ExpectQuery("Select name from users").WillReturnError(errInternalConn)

	_, executionErr := service.GetNames()

	require.Error(t, executionErr)
	assert.Contains(t, executionErr.Error(), "db query")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_ScanError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("Select name from users").WillReturnRows(rows)

	_, executionErr := service.GetNames()

	require.Error(t, executionErr)
	assert.Contains(t, executionErr.Error(), "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNames_RowsIterationError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Yana").RowError(0, errScanFailed)
	mock.ExpectQuery("Select name from users").WillReturnRows(rows)

	_, executionErr := service.GetNames()

	require.Error(t, executionErr)
	assert.Contains(t, executionErr.Error(), "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_SuccessCase(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Charlie")
	mock.ExpectQuery("Select distinct name from users").WillReturnRows(rows)

	names, executionErr := service.GetUniqueNames()

	require.NoError(t, executionErr)
	assert.Equal(t, []string{"Charlie"}, names)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_QueryError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)

	mock.ExpectQuery("Select distinct name from users").WillReturnError(errFatalDB)

	_, executionErr := service.GetUniqueNames()

	require.Error(t, executionErr)
	assert.Contains(t, executionErr.Error(), "db query")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_ScanError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	mock.ExpectQuery("Select distinct name from users").WillReturnRows(rows)

	_, executionErr := service.GetUniqueNames()

	require.Error(t, executionErr)
	assert.Contains(t, executionErr.Error(), "rows scanning")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUniqueNames_RowsIterationError(t *testing.T) {
	t.Parallel()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)
	rows := sqlmock.NewRows([]string{"name"}).AddRow("Charlie").RowError(0, errInterrupted)
	mock.ExpectQuery("Select distinct name from users").WillReturnRows(rows)

	_, executionErr := service.GetUniqueNames()

	require.Error(t, executionErr)
	assert.Contains(t, executionErr.Error(), "rows error")
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestNewService(t *testing.T) {
	t.Parallel()

	sqlDB, _, err := sqlmock.New()
	require.NoError(t, err)

	defer func() { _ = sqlDB.Close() }()

	service := dbservice.New(sqlDB)
	assert.NotNil(t, service.DB)
}
