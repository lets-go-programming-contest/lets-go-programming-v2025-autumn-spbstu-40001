package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/verticalochka/task-6/internal/db"
)

func TestGetActiveUsers(t *testing.T) {
	t.Parallel()

	dbMock, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbMock.Close()

	service := db.New(dbMock)

	rows := sqlmock.NewRows([]string{"username"}).
		AddRow("alice").
		AddRow("bob")

	mock.ExpectQuery("SELECT username FROM users WHERE active = true").
		WillReturnRows(rows)

	result, err := service.GetActiveUsers()
	require.NoError(t, err)
	require.Equal(t, []string{"alice", "bob"}, result)

	mock.ExpectQuery("SELECT username FROM users WHERE active = true").
		WillReturnError(errors.New("connection failed"))

	result, err = service.GetActiveUsers()
	require.Error(t, err)
	require.ErrorContains(t, err, "query error")
	require.Nil(t, result)
}

func TestDeactivateUser(t *testing.T) {
	t.Parallel()

	dbMock, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer dbMock.Close()

	service := db.New(dbMock)

	mock.ExpectExec("UPDATE users SET active = false WHERE username = ?").
		WithArgs("inactive_user").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = service.DeactivateUser("inactive_user")
	require.NoError(t, err)

	mock.ExpectExec("UPDATE users SET active = false WHERE username = ?").
		WithArgs("nonexistent").
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = service.DeactivateUser("nonexistent")
	require.Error(t, err)
	require.ErrorContains(t, err, "user not found")
}
