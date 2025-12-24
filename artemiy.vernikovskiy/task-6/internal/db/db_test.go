package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	taskDbPack "github.com/Aapng-cmd/task-6/internal/db"
)

var (
	ErrQueryError   = errors.New("we got the news for you, you are doomed")
	ErrRowsScanning = errors.New("rows scanning")
	ErrRowsError    = errors.New("rows error")
	ErrDBQuery      = errors.New("db query")
)

const (
	sqlGetNames       = "SELECT name FROM users"
	sqlGetUniqueNames = "SELECT DISTINCT name FROM users"
)

func TestDBServiceGetNamesSuccess(t *testing.T) {
	const numberOfData = 3

	testNames := []string{"Petya", "Vanya", "Punk"}

	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range testNames {
		rows.AddRow(name)
	}

	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Len(t, names, numberOfData)

	assert.Equal(t, testNames, names)
}

func TestDBServiceGetNamesEmpty(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestDBServiceGetNamesScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	testNames := []interface{}{"Fuagra", nil, "Fukh"}

	rows := sqlmock.NewRows([]string{"name"})

	for _, name := range testNames {
		rows.AddRow(name)
	}

	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), ErrRowsScanning.Error())
}

func TestDBServiceGetNamesRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Pihta")
	rows.RowError(0, sql.ErrTxDone)

	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), ErrRowsError.Error())
}

func TestDBServiceGetNamesQueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	mock.ExpectQuery(sqlGetNames).WillReturnError(ErrQueryError)

	service := taskDbPack.New(mockDB)
	names, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), ErrDBQuery.Error())
}

func TestDBServiceGetUniqueNamesSuccess(t *testing.T) {
	const numberOfData = 3

	testNames := []string{"UniqueName1FantasyDied", "AgroCultureIsTheBest", "GodLovesNumber3Large"} // not a blasphemy

	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range testNames {
		rows.AddRow(name)
	}

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	assert.Len(t, names, numberOfData)

	assert.Equal(t, testNames, names)
}

func TestDBServiceGetUniqueNamesEmpty(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestDBServiceGetUniqueNamesScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	testNames := []interface{}{"GetReadyForNil", nil, "Gotcha"}

	rows := sqlmock.NewRows([]string{"name"})

	for _, name := range testNames {
		rows.AddRow(name)
	}

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), ErrRowsScanning.Error())
}

func TestDBServiceGetUniqueNamesRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("OhohSneaky")
	rows.RowError(0, sql.ErrTxDone)

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := taskDbPack.New(mockDB)
	names, err := service.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), ErrRowsError.Error())
}

func TestDBServiceGetUniqueNamesQueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnError(ErrQueryError)

	service := taskDbPack.New(mockDB)
	names, err := service.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), ErrDBQuery.Error())
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	service := taskDbPack.New(mockDB)
	assert.Equal(t, mockDB, service.DB)
}
