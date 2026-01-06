package db_test

import (
	"testing"

	"github.com/A1exCRE/task-6/internal/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	dbConn, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbConn.Close()

	svc := db.New(dbConn)
	assert.Equal(t, dbConn, svc.DB)
}

func TestGetNames(t *testing.T) {
	dbConn, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbConn.Close()

	svc := db.DBService{DB: dbConn}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alex").
		AddRow("Maria")

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := svc.GetNames()

	assert.NoError(t, err)
	assert.Equal(t, []string{"Alex", "Maria"}, names)
	assert.NoError(t, mock.ExpectationsWereMet())
}
