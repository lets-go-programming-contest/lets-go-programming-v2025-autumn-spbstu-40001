package db_tests

import (
	"database/sql"
	"testing"

	"github.com/Aapng-cmd/task-6/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	sqlGetNames       = "SELECT name FROM users"
	sqlGetUniqueNames = "SELECT DISTINCT name FROM users"
)

func TestDBServiceGetNamesSuccess(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Petya").
		AddRow("Vanya").
		AddRow("Punk")

	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Len(t, names, 3)
	assert.Equal(t, "Petya", names[0])
	assert.Equal(t, "Vanya", names[1])
	assert.Equal(t, "Punk", names[2])
}

func TestDBServiceGetNamesEmpty(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetNames()

	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestDBServiceGetNamesScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Fuagra").
		AddRow(nil).
		AddRow("Fukh")

	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows scanning")
}

func TestDBServiceGetNamesRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	rows.AddRow("Pihta")
	rows.RowError(0, sql.ErrTxDone)

	mock.ExpectQuery(sqlGetNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows error")
}

func TestDBServiceGetUniqueNamesSuccess(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("UniqueName1FantasyDied").
		AddRow("AgroCultureIsTheBest").
		AddRow("GodLovesNumber3Large") // not a blasphemy

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	assert.Len(t, names, 3)
	assert.Equal(t, "UniqueName1FantasyDied", names[0])
	assert.Equal(t, "AgroCultureIsTheBest", names[1])
	assert.Equal(t, "GodLovesNumber3Large", names[2])
}

func TestDBServiceGetUniqueNamesEmpty(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetUniqueNames()

	require.NoError(t, err)
	assert.Empty(t, names)
}

func TestDBServiceGetUniqueNamesQueryError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnError(sql.ErrConnDone)

	service := db.New(mockDB)
	names, err := service.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "db query")
}

func TestDBServiceGetUniqueNamesScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("GetReadyForNil").
		AddRow(nil).
		AddRow("Gotcha")

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows scanning")
}

func TestDBServiceGetUniqueNamesRowsError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	rows := sqlmock.NewRows([]string{"name"})
	rows.AddRow("OhohSneaky")
	rows.RowError(0, sql.ErrTxDone)

	mock.ExpectQuery(sqlGetUniqueNames).WillReturnRows(rows)

	service := db.New(mockDB)
	names, err := service.GetUniqueNames()

	require.Error(t, err)
	assert.Nil(t, names)
	assert.Contains(t, err.Error(), "rows error")
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB, _, err := sqlmock.New()
	require.NoError(t, err)

	defer mockDB.Close()

	service := db.New(mockDB)
	assert.Equal(t, mockDB, service.DB)
}
