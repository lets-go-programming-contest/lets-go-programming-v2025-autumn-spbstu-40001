package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/DimasFantomasA/task-6/internal/db"
)

var (
	errConnectionFailed = errors.New("connection failed")
	errRow              = errors.New("row error")
)

type DBServiceTestSuite struct {
	suite.Suite
	mockDB *sql.DB
	mock   sqlmock.Sqlmock
}

func (s *DBServiceTestSuite) SetupTest() {
	var err error
	s.mockDB, s.mock, err = sqlmock.New()
	s.Require().NoError(err)
}

func (s *DBServiceTestSuite) TearDownTest() {
	s.mockDB.Close()
}

func (s *DBServiceTestSuite) TestNew() {
	service := db.New(s.mockDB)
	s.Equal(s.mockDB, service.DB)
}

func (s *DBServiceTestSuite) TestGetNames_Success() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Petr").
		AddRow("Anna")

	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	s.Require().NoError(err)
	s.Equal([]string{"Ivan", "Petr", "Anna"}, result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_EmptyResult() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"})
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	s.Require().NoError(err)
	s.Empty(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_QueryError() {
	service := db.DBService{DB: s.mockDB}

	testError := errConnectionFailed
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnError(testError)

	result, err := service.GetNames()

	s.Require().Error(err)
	s.Require().ErrorContains(err, "db query")
	s.Contains(err.Error(), testError.Error())
	s.Nil(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_ScanError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	s.Require().Error(err)
	s.Require().ErrorContains(err, "rows scanning")
	s.Nil(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetNames_RowsError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errRow)
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	s.Require().Error(err)
	s.Require().ErrorContains(err, "rows error")
	s.Nil(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_Success() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		AddRow("Petr").
		AddRow("Anna")

	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	s.Require().NoError(err)
	s.Equal([]string{"Ivan", "Petr", "Anna"}, result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_EmptyResult() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"})
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	s.Require().NoError(err)
	s.Empty(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_QueryError() {
	service := db.DBService{DB: s.mockDB}

	testError := errConnectionFailed
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(testError)

	result, err := service.GetUniqueNames()

	s.Require().Error(err)
	s.Require().ErrorContains(err, "db query")
	s.Contains(err.Error(), testError.Error())
	s.Nil(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_ScanError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	s.Require().Error(err)
	s.Require().ErrorContains(err, "rows scanning")
	s.Nil(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func (s *DBServiceTestSuite) TestGetUniqueNames_RowsError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Ivan").
		RowError(0, errRow)
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	s.Require().Error(err)
	s.Require().ErrorContains(err, "rows error")
	s.Nil(result)
	s.Require().NoError(s.mock.ExpectationsWereMet())
}

func TestDBServiceTestSuite(t *testing.T) {
	t.Parallel() // Добавлена эта строка
	suite.Run(t, new(DBServiceTestSuite))
}
