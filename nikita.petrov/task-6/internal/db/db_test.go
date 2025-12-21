package db_test

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nekich06/task-6/internal/db"
	"github.com/stretchr/testify/suite"
)

const (
	testQuerySelectNames       = "SELECT name FROM users"
	testQuerySelectUniqueNames = "SELECT DISTINCT name FROM users"
)

var (
	testNames       = []string{"Alice", "Bob", "Charlie", "Alice", "Bob"}
	testUniqueNames = []string{"Alice", "Bob", "Charlie"}
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

func TestDBServiceTestSuite(t *testing.T) {
	suite.Run(t, new(DBServiceTestSuite))
}

func (s *DBServiceTestSuite) TestNew() {
	service := db.New(s.mockDB)
	s.NotNil(service)
	s.Equal(s.mockDB, service.DB)
}

func (s *DBServiceTestSuite) TestGetNames_SuccessWithData() {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range testNames {
		rows.AddRow(name)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectNames)).WillReturnRows(rows)

	names, err := s.service.GetNames()

	s.Require().NoError(err)
	s.Equal(testNames, names)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_EmptyResult() {
	rows := sqlmock.NewRows([]string{"name"})

	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectNames)).WillReturnRows(rows)

	names, err := s.service.GetNames()

	s.Require().NoError(err)
	s.Empty(names)
	s.Len(names, 0)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_DBError() {
	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectNames)).
		WillReturnError(errDB)

	names, err := s.service.GetNames()

	s.Require().Error(err)
	s.Nil(names)
	s.Contains(err.Error(), "db query")
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_RowsScanningError() {
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow(nil).
		RowError(1, errors.New("scan error"))

	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectNames)).
		WillReturnRows(rows)

	names, err := s.service.GetNames()

	s.Require().Error(err)
	s.Nil(names)
	s.Contains(err.Error(), "rows scanning")
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_RowsIterationError() {
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		AddRow("Charlie").
		CloseError(errRowIteration)

	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectNames)).
		WillReturnRows(rows)

	names, err := s.service.GetNames()

	s.Require().Error(err)
	s.Nil(names)
	s.Contains(err.Error(), "rows error")
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_SuccessWithData() {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range testUniqueNames {
		rows.AddRow(name)
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectUniqueNames)).
		WillReturnRows(rows)

	names, err := s.service.GetUniqueNames()

	s.Require().NoError(err)
	s.Equal(testUniqueNames, names)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_EmptyResult() {
	rows := sqlmock.NewRows([]string{"name"})

	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectUniqueNames)).
		WillReturnRows(rows)

	names, err := s.service.GetUniqueNames()

	s.Require().NoError(err)
	s.Empty(names)
	s.Len(names, 0)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_DBError() {
	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectUniqueNames)).
		WillReturnError(errDB)

	names, err := s.service.GetUniqueNames()

	s.Require().Error(err)
	s.Nil(names)
	s.Contains(err.Error(), "db query")
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_WithDuplicatesInDB() {
	dbRows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob").
		AddRow("Alice").
		AddRow("Charlie").
		AddRow("Bob")

	expectedUnique := []string{"Alice", "Bob", "Charlie"}

	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectUniqueNames)).
		WillReturnRows(dbRows)

	names, err := s.service.GetUniqueNames()

	s.Require().NoError(err)
	s.Equal(expectedUnique, names)
	s.NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestServiceMethodsUseConstants() {
	rows1 := sqlmock.NewRows([]string{"name"}).AddRow("test")
	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectNames)).WillReturnRows(rows1)

	_, err := s.service.GetNames()
	s.Require().NoError(err)

	rows2 := sqlmock.NewRows([]string{"name"}).AddRow("test")
	s.mock.ExpectQuery(regexp.QuoteMeta(testQuerySelectUniqueNames)).WillReturnRows(rows2)

	_, err = s.service.GetUniqueNames()
	s.Require().NoError(err)

	s.NoError(s.mock.ExpectationsWereMet())
}
