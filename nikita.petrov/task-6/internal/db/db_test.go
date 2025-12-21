package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Nekich06/task-6/internal/db"
	"github.com/stretchr/testify/suite"
)

type DataServiceTestSuite struct {
	suite.Suite
	dbConnection *sql.DB
	sqlMock      sqlmock.Sqlmock
}

func (s *DataServiceTestSuite) SetupSuite() {
	var setupErr error
	s.dbConnection, s.sqlMock, setupErr = sqlmock.New()
	s.Require().Nil(setupErr)
}

func (s *DataServiceTestSuite) TearDownSuite() {
	if s.dbConnection != nil {
		s.dbConnection.Close()
	}
}

func (s *DataServiceTestSuite) TestConstructor() {
	dataService := db.New(s.dbConnection)
	s.Equal(s.dbConnection, dataService.DB)
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

	s.Nil(fetchErr)
	s.Equal(expectedData, actualResult)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsersEmptyDataset() {
	dataHandler := db.DBService{DB: s.dbConnection}

	emptyRows := sqlmock.NewRows([]string{"name"})
	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(emptyRows)

	resultData, fetchErr := dataHandler.GetNames()

	s.Nil(fetchErr)
	s.Empty(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsersDatabaseFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}

	connectionFailure := errors.New("database unreachable")
	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnError(connectionFailure)

	resultData, fetchErr := dataHandler.GetNames()

	s.NotNil(fetchErr)
	s.ErrorContains(fetchErr, "db query")
	s.Contains(fetchErr.Error(), connectionFailure.Error())
	s.Nil(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsersRowParsingFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}

	faultyRows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(faultyRows)

	resultData, fetchErr := dataHandler.GetNames()

	s.NotNil(fetchErr)
	s.ErrorContains(fetchErr, "rows scanning")
	s.Nil(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestFetchAllUsersRowIterationFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}

	problematicRows := sqlmock.NewRows([]string{"name"}).AddRow("Michael").RowError(0, errors.New("iterator broken"))
	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(problematicRows)

	resultData, fetchErr := dataHandler.GetNames()

	s.NotNil(fetchErr)
	s.ErrorContains(fetchErr, "rows error")
	s.Nil(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
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

	s.Nil(fetchErr)
	s.Equal(uniqueData, actualResult)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsersEmptyDataset() {
	dataHandler := db.DBService{DB: s.dbConnection}

	emptyRows := sqlmock.NewRows([]string{"name"})
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(emptyRows)

	resultData, fetchErr := dataHandler.GetUniqueNames()

	s.Nil(fetchErr)
	s.Empty(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsersDatabaseFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}

	connectionFailure := errors.New("query execution failed")
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnError(connectionFailure)

	resultData, fetchErr := dataHandler.GetUniqueNames()

	s.NotNil(fetchErr)
	s.ErrorContains(fetchErr, "db query")
	s.Contains(fetchErr.Error(), connectionFailure.Error())
	s.Nil(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsersRowParsingFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}

	faultyRows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(faultyRows)

	resultData, fetchErr := dataHandler.GetUniqueNames()

	s.NotNil(fetchErr)
	s.ErrorContains(fetchErr, "rows scanning")
	s.Nil(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsersRowIterationFailure() {
	dataHandler := db.DBService{DB: s.dbConnection}

	problematicRows := sqlmock.NewRows([]string{"name"}).AddRow("Elizabeth").RowError(0, errors.New("iterator malfunction"))
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(problematicRows)

	resultData, fetchErr := dataHandler.GetUniqueNames()

	s.NotNil(fetchErr)
	s.ErrorContains(fetchErr, "rows error")
	s.Nil(resultData)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestRetrieveDistinctUsersDuplicateFiltering() {
	dataHandler := db.DBService{DB: s.dbConnection}

	uniqueEntries := []string{"Benjamin", "Charlotte"}
	mockRows := sqlmock.NewRows([]string{"name"})

	for _, entry := range uniqueEntries {
		mockRows.AddRow(entry)
	}

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetUniqueNames()

	s.Nil(fetchErr)
	s.Equal(uniqueEntries, actualResult)
	s.Len(actualResult, 2)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestServiceHandlesMultipleInvocations() {
	dataHandler := db.DBService{DB: s.dbConnection}

	firstRows := sqlmock.NewRows([]string{"name"}).AddRow("Thomas")
	secondRows := sqlmock.NewRows([]string{"name"}).AddRow("Emma")

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(firstRows)
	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(secondRows)

	firstResult, firstErr := dataHandler.GetNames()
	s.Nil(firstErr)
	s.Equal([]string{"Thomas"}, firstResult)

	secondResult, secondErr := dataHandler.GetUniqueNames()
	s.Nil(secondErr)
	s.Equal([]string{"Emma"}, secondResult)

	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestServiceWithInvalidConnection() {
	brokenConnection, _, _ := sqlmock.New()
	brokenConnection.Close()

	dataHandler := db.DBService{DB: brokenConnection}

	_, fetchErr := dataHandler.GetNames()
	s.NotNil(fetchErr)
}

func (s *DataServiceTestSuite) TestServiceWithSpecialCharacters() {
	dataHandler := db.DBService{DB: s.dbConnection}

	testData := []string{"José", "Renée", "Björn", "Siobhán"}
	mockRows := sqlmock.NewRows([]string{"name"})

	for _, item := range testData {
		mockRows.AddRow(item)
	}

	s.sqlMock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetNames()

	s.Nil(fetchErr)
	s.Equal(testData, actualResult)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func (s *DataServiceTestSuite) TestServiceWithMixedCaseData() {
	dataHandler := db.DBService{DB: s.dbConnection}

	testData := []string{"alex", "ALEX", "Alex", "aLeX"}
	mockRows := sqlmock.NewRows([]string{"name"})

	for _, item := range testData {
		mockRows.AddRow(item)
	}

	s.sqlMock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockRows)

	actualResult, fetchErr := dataHandler.GetUniqueNames()

	s.Nil(fetchErr)
	s.Equal(testData, actualResult)
	s.Nil(s.sqlMock.ExpectationsWereMet())
}

func TestDataServiceTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(DataServiceTestSuite))
}
