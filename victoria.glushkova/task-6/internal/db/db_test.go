package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vikaglushkova/task-6/internal/db"
	"github.com/stretchr/testify/suite"
)

var (
	errUnreachableDB  = errors.New("database is unreachable")
	errBadQueryExec   = errors.New("query execution failed")
	errBrokenIterator = errors.New("iterator broken")
	errClosedRows     = errors.New("rows closed")
)

type DataServiceTestSuite struct {
	suite.Suite
	dbConnection *sql.DB
	sqlMock      sqlmock.Sqlmock
}

func (s *DataServiceTestSuite) SetupSuite() {
	var setupErr error
	s.dbConnection, s.sqlMock, setupErr = sqlmock.New()
	s.Require().NoError(setupErr)
}

func (s *DataServiceTestSuite) TearDownSuite() {
	if s.dbConnection != nil {
		s.dbConnection.Close()
	}
}

func (s *DataServiceTestSuite) TestConstructor() {
	dataService := db.New(s.dbConnection)
	s.Require().Equal(s.dbConnection, dataService.DB)
}

func (s *DataServiceTestSuite) TestFetchAllUsers() {
	dataHandler := db.DBService{DB: s.dbConnection}
	expectedData := []string{"Michael", "Sarah", "William"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, item := range expectedData {
		mockRows.AddRow(item)
	}

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetNames()
	s.Require().NoError(fetchErr)
	s.Require().Equal(expectedData, actualResult)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsers_EmptyDataset() {
	dataHandler := db.DBService{DB: s.dbConnection}
	emptyRows := sqlmock.NewRows([]string{"name"})

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(emptyRows)

	resultData, fetchErr := dataHandler.GetNames()
	s.Require().NoError(fetchErr)
	s.Require().Empty(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsers_DatabaseFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}
	connectionFailure := errUnreachableDB

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnError(connectionFailure)

	resultData, fetchErr := dataHandler.GetNames()
	s.Require().Error(fetchErr)
	s.Require().Contains(fetchErr.Error(), "db query")
	s.Require().Contains(fetchErr.Error(), connectionFailure.Error())
	s.Require().Nil(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsers_RowParsingFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}
	faultyRows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(faultyRows)

	resultData, fetchErr := dataHandler.GetNames()
	s.Require().Error(fetchErr)
	s.Require().Contains(fetchErr.Error(), "rows scanning")
	s.Require().Nil(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsers_RowIterationFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}
	problematicRows := sqlmock.NewRows([]string{"name"}).AddRow("Michael").RowError(0, errBrokenIterator)

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(problematicRows)

	resultData, fetchErr := dataHandler.GetNames()
	s.Require().Error(fetchErr)
	s.Require().Contains(fetchErr.Error(), "rows error")
	s.Require().Contains(fetchErr.Error(), errBrokenIterator.Error())
	s.Require().Nil(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsers() {
	dataHandler := db.DBService{DB: s.dbConnection}
	uniqueData := []string{"Elizabeth", "James", "Olivia"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, item := range uniqueData {
		mockRows.AddRow(item)
	}

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetUniqueNames()
	s.Require().NoError(fetchErr)
	s.Require().Equal(uniqueData, actualResult)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsers_EmptyDataset() {
	dataHandler := db.DBService{DB: s.dbConnection}
	emptyRows := sqlmock.NewRows([]string{"name"})

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(emptyRows)

	resultData, fetchErr := dataHandler.GetUniqueNames()
	s.Require().NoError(fetchErr)
	s.Require().Empty(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsers_DatabaseFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}
	connectionFailure := errBadQueryExec

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(connectionFailure)

	resultData, fetchErr := dataHandler.GetUniqueNames()
	s.Require().Error(fetchErr)
	s.Require().Contains(fetchErr.Error(), "db query")
	s.Require().Contains(fetchErr.Error(), connectionFailure.Error())
	s.Require().Nil(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsers_RowParsingFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}
	faultyRows := sqlmock.NewRows([]string{"name"}).AddRow(nil)

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(faultyRows)

	resultData, fetchErr := dataHandler.GetUniqueNames()
	s.Require().Error(fetchErr)
	s.Require().Contains(fetchErr.Error(), "rows scanning")
	s.Require().Nil(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsers_RowIterationFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}
	problematicRows := sqlmock.NewRows([]string{"name"}).AddRow("Elizabeth").RowError(0, errBrokenIterator)

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(problematicRows)

	resultData, fetchErr := dataHandler.GetUniqueNames()
	s.Require().Error(fetchErr)
	s.Require().Contains(fetchErr.Error(), "rows error")
	s.Require().Contains(fetchErr.Error(), errBrokenIterator.Error())
	s.Require().Nil(resultData)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsers_DuplicateFiltering() {
	dataHandler := db.DBService{DB: s.dbConnection}
	uniqueEntries := []string{"Benjamin", "Charlotte"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, entry := range uniqueEntries {
		mockRows.AddRow(entry)
	}

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetUniqueNames()
	s.Require().NoError(fetchErr)
	s.Require().Equal(uniqueEntries, actualResult)
	s.Require().Len(actualResult, 2)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestService_HandlesMultipleInvocations() {
	dataHandler := db.DBService{DB: s.dbConnection}
	firstRows := sqlmock.NewRows([]string{"name"}).AddRow("Thomas")
	secondRows := sqlmock.NewRows([]string{"name"}).AddRow("Emma")

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(firstRows)
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(secondRows)

	firstResult, firstErr := dataHandler.GetNames()
	s.Require().NoError(firstErr)
	s.Require().Equal([]string{"Thomas"}, firstResult)

	secondResult, secondErr := dataHandler.GetUniqueNames()
	s.Require().NoError(secondErr)
	s.Require().Equal([]string{"Emma"}, secondResult)

	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestService_WithSpecialCharacters() {
	dataHandler := db.DBService{DB: s.dbConnection}
	testData := []string{"José", "Renée", "Björn", "Siobhán"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, item := range testData {
		mockRows.AddRow(item)
	}

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetNames()
	s.Require().NoError(fetchErr)
	s.Require().Equal(testData, actualResult)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestService_WithMixedCaseData() {
	dataHandler := db.DBService{DB: s.dbConnection}
	testData := []string{"alex", "ALEX", "Alex", "aLeX"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, item := range testData {
		mockRows.AddRow(item)
	}

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetUniqueNames()
	s.Require().NoError(fetchErr)
	s.Require().Equal(testData, actualResult)
	s.Require().NoError(s.sqlMock.ExpectationsWereMet())
}

func TestDataServiceTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(DataServiceTestSuite))
}
