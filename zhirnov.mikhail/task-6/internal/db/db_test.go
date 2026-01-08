package db_test

import (
	"errors"
	"testing"

	"github.com/m1shanyanba/task-6/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errQuery = errors.New("query failed")
	errRows  = errors.New("rows error")
)

func TestGetNames(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).AddRow("Polina").AddRow("Anya")
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		service := db.New(sqlDB)
		names, err := service.GetNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"Polina", "Anya"}, names)
	})

	t.Run("query_error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		mock.ExpectQuery("SELECT name FROM users").WillReturnError(errQuery)

		service := db.New(sqlDB)
		_, err = service.GetNames()

		require.ErrorIs(t, err, errQuery)
	})

	t.Run("scan_error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		service := db.New(sqlDB)
		_, err = service.GetNames()

		require.Error(t, err)
	})

	t.Run("rows_error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).AddRow("Polina").CloseError(errRows)
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

		service := db.New(sqlDB)
		_, err = service.GetNames()

		require.ErrorIs(t, err, errRows)
	})
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).AddRow("Polina").AddRow("Anya")
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

		service := db.New(sqlDB)
		names, err := service.GetUniqueNames()

		require.NoError(t, err)
		assert.Equal(t, []string{"Polina", "Anya"}, names)
	})

	t.Run("query_error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(errQuery)

		service := db.New(sqlDB)
		_, err = service.GetUniqueNames()

		require.ErrorIs(t, err, errQuery)
	})

	t.Run("scan_error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

		service := db.New(sqlDB)
		_, err = service.GetUniqueNames()

		require.Error(t, err)
	})

	t.Run("rows_error", func(t *testing.T) {
		t.Parallel()

		sqlDB, mock, err := sqlmock.New()
		require.NoError(t, err)
		defer sqlDB.Close()

		rows := sqlmock.NewRows([]string{"name"}).AddRow("Polina").CloseError(errRows)
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

		service := db.New(sqlDB)
		_, err = service.GetUniqueNames()

		require.ErrorIs(t, err, errRows)
	})
}
