package db_test

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nekich06/task-6/internal/db"
	"github.com/stretchr/testify/suite"
)

var (
	errDB           = errors.New("db error")
	errRowIteration = errors.New("row iteration error")
)

type DBServiceTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mock    sqlmock.Sqlmock
	service db.DBService
}

func (s *DBServiceTestSuite) SetupTest() {
	var err error
	s.mockDB, s.mock, err = sqlmock.New()
	s.Require().NoError(err, "failed to create sqlmock")
	s.service = db.New(s.mockDB)
}

func (s *DBServiceTestSuite) TearDownTest() {
	if s.mockDB != nil {
		s.mockDB.Close()
	}
}

func (s *DBServiceTestSuite) TestNew() {
	service := db.New(s.mockDB)
	s.NotNil(service)
	s.Equal(s.mockDB, service.DB)
}

func (s *DBServiceTestSuite) TestGetNames_SuccessWithData() {
	expectedNames := []string{"Alice", "Bob", "Charlie"}
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range expectedNames {
		rows.AddRow(name)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT name FROM users")).WillReturnRows(rows)

	names, err := s.service.GetNames()

	s.Require().NoError(err)
	s.Equal(expectedNames, names)
	s.NoError(s.mock.ExpectationsWereMet())
}
