package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/netwite/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DBServiceTestSuite struct {
	suite.Suite
	mockDB *sql.DB
	mock   sqlmock.Sqlmock
}

func (s *DBServiceTestSuite) SetupTest() {
	var err error
	s.mockDB, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Fatalf("failed to create sqlmock: %v", err)
	}
}

func (s *DBServiceTestSuite) TearDownTest() {
	if s.mockDB != nil {
		s.mockDB.Close()
	}
}

func (s *DBServiceTestSuite) TestNew() {
	service := db.New(s.mockDB)
	assert.Equal(s.T(), s.mockDB, service.DB)
}

func (s *DBServiceTestSuite) TestGetNames_Success() {
	service := db.DBService{DB: s.mockDB}

	expectedRows := []string{"Alice", "Bob", "Charlie"}
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range expectedRows {
		rows.AddRow(name)
	}

	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), expectedRows, result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetNames_EmptyResult() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"})
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	assert.NoError(s.T(), err)
	assert.Empty(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetNames_QueryError() {
	service := db.DBService{DB: s.mockDB}

	testError := errors.New("connection failed")
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnError(testError)

	result, err := service.GetNames()

	assert.Error(s.T(), err)
	assert.ErrorContains(s.T(), err, "db query")
	assert.Contains(s.T(), err.Error(), testError.Error())
	assert.Nil(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetNames_ScanError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	assert.Error(s.T(), err)
	assert.ErrorContains(s.T(), err, "rows scanning")
	assert.Nil(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetNames_RowsError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(0, errors.New("row error"))
	s.mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	result, err := service.GetNames()

	assert.Error(s.T(), err)
	assert.ErrorContains(s.T(), err, "rows error")
	assert.Nil(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetUniqueNames_Success() {
	service := db.DBService{DB: s.mockDB}

	uniqueNames := []string{"Alice", "Bob", "Charlie"}

	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range uniqueNames {
		rows.AddRow(name)
	}

	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), uniqueNames, result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetUniqueNames_EmptyResult() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"})
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	assert.NoError(s.T(), err)
	assert.Empty(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetUniqueNames_QueryError() {
	service := db.DBService{DB: s.mockDB}

	testError := errors.New("connection failed")
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(testError)

	result, err := service.GetUniqueNames()

	assert.Error(s.T(), err)
	assert.ErrorContains(s.T(), err, "db query")
	assert.Contains(s.T(), err.Error(), testError.Error())
	assert.Nil(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetUniqueNames_ScanError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	assert.Error(s.T(), err)
	assert.ErrorContains(s.T(), err, "rows scanning")
	assert.Nil(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetUniqueNames_RowsError() {
	service := db.DBService{DB: s.mockDB}

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Alice").RowError(0, errors.New("row error"))
	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	assert.Error(s.T(), err)
	assert.ErrorContains(s.T(), err, "rows error")
	assert.Nil(s.T(), result)

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func (s *DBServiceTestSuite) TestGetUniqueNames_ReturnsOnlyUnique() {
	service := db.DBService{DB: s.mockDB}

	uniqueRows := []string{"John", "Jane"}
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range uniqueRows {
		rows.AddRow(name)
	}

	s.mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(rows)

	result, err := service.GetUniqueNames()

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), uniqueRows, result)
	assert.Equal(s.T(), 2, len(result), "Должно вернуть только уникальные значения")

	err = s.mock.ExpectationsWereMet()
	assert.NoError(s.T(), err)
}

func TestDBServiceTestSuite(t *testing.T) {
	suite.Run(t, new(DBServiceTestSuite))
}
