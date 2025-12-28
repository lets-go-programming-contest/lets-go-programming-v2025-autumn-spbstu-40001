package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/vikaglushkova/task-6/internal/db"
)

type DataServiceTestSuite struct {
	suite.Suite
	dbConnection *sql.DB
	sqlMock      sqlmock.Sqlmock
}

func (s *DataServiceTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	s.Require().NoError(err)
	s.dbConnection = conn
	s.sqlMock = mock
}

func (s *DataServiceTestSuite) TearDownTest() {
	s.dbConnection.Close()
}

func (s *DataServiceTestSuite) TestConstructor() {
	service := db.New(s.dbConnection)
	s.Require().Equal(s.dbConnection, service.DB)
}

func (s *DataServiceTestSuite) TestGetNames_Success() {
	service := db.New(s.dbConnection)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")

	s.sqlMock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetNames()
	s.Require().NoError(err)
	s.Require().Equal([]string{"Alice", "Bob"}, result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetNames_Empty() {
	service := db.New(s.dbConnection)

	rows := sqlmock.NewRows([]string{"name"})

	s.sqlMock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetNames()
	s.Require().NoError(err)
	s.Require().Empty(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetNames_DatabaseError() {
	service := db.New(s.dbConnection)

	dbErr := errors.New("database error")
	s.sqlMock.ExpectQuery("SELECT name FROM users").
		WillReturnError(dbErr)

	result, err := service.GetNames()
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "db query")
	s.Require().Nil(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetNames_RowScanError() {
	service := db.New(s.dbConnection)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	s.sqlMock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetNames()
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "rows scanning")
	s.Require().Nil(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetNames_RowIterationError() {
	service := db.New(s.dbConnection)

	rowErr := errors.New("row error")
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		RowError(0, rowErr)

	s.sqlMock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetNames()
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "rows error")
	s.Require().Nil(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetUniqueNames_Success() {
	service := db.New(s.dbConnection)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetUniqueNames()
	s.Require().NoError(err)
	s.Require().Equal([]string{"Alice", "Bob"}, result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetUniqueNames_Empty() {
	service := db.New(s.dbConnection)

	rows := sqlmock.NewRows([]string{"name"})

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetUniqueNames()
	s.Require().NoError(err)
	s.Require().Empty(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetUniqueNames_DatabaseError() {
	service := db.New(s.dbConnection)

	dbErr := errors.New("database error")
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(dbErr)

	result, err := service.GetUniqueNames()
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "db query")
	s.Require().Nil(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetUniqueNames_RowScanError() {
	service := db.New(s.dbConnection)

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetUniqueNames()
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "rows scanning")
	s.Require().Nil(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestGetUniqueNames_RowIterationError() {
	service := db.New(s.dbConnection)

	rowErr := errors.New("row error")
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		RowError(0, rowErr)

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetUniqueNames()
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "rows error")
	s.Require().Nil(result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestMultipleMethodCalls() {
	service := db.New(s.dbConnection)

	firstRows := sqlmock.NewRows([]string{"name"}).AddRow("First")
	secondRows := sqlmock.NewRows([]string{"name"}).AddRow("Second")

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(firstRows)
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(secondRows)

	firstResult, firstErr := service.GetNames()
	s.Require().NoError(firstErr)
	s.Require().Equal([]string{"First"}, firstResult)

	secondResult, secondErr := service.GetUniqueNames()
	s.Require().NoError(secondErr)
	s.Require().Equal([]string{"Second"}, secondResult)

	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestServiceWithSpecialCharacters() {
	service := db.New(s.dbConnection)

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("José").
		AddRow("Müller").
		AddRow("Françoise")

	s.sqlMock.ExpectQuery("SELECT name FROM users").
		WillReturnRows(rows)

	result, err := service.GetNames()
	s.Require().NoError(err)
	s.Require().Equal([]string{"José", "Müller", "Françoise"}, result)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func TestDataServiceTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(DataServiceTestSuite))
}
