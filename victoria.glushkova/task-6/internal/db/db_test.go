package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/vikaglushkova/task-6/internal/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DatabaseOperationsTestSuite struct {
	suite.Suite
	dbConn       *sql.DB
	mockExecutor sqlmock.Sqlmock
	dataService  db.DBService
}

func (suite *DatabaseOperationsTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	require.NoError(suite.T(), err, "Mock DB creation failed")
	suite.dbConn = conn
	suite.mockExecutor = mock
	suite.dataService = db.New(conn)
}

func (suite *DatabaseOperationsTestSuite) TearDownTest() {
	if suite.dbConn != nil {
		suite.dbConn.Close()
	}
}

func (suite *DatabaseOperationsTestSuite) TestConstructorNew() {
	conn, mock, err := sqlmock.New()
	require.NoError(suite.T(), err)
	defer conn.Close()

	service := db.New(conn)
	assert.NotNil(suite.T(), service)
	assert.Equal(suite.T(), conn, service.DB)

	assert.NoError(suite.T(), mock.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestServiceInitialization() {
	dbConn := suite.dataService.DB
	assert.NotNil(suite.T(), dbConn, "Database connection should be established")
}

func (suite *DatabaseOperationsTestSuite) TestFetchUserListSuccess() {
	expectedUsers := []string{"Alex", "Maria", "John"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, user := range expectedUsers {
		mockRows.AddRow(user)
	}

	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnRows(mockRows)

	result, err := suite.dataService.GetNames()

	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUsers, result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchUserListEmptyResult() {
	emptyRows := sqlmock.NewRows([]string{"name"})
	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnRows(emptyRows)

	result, err := suite.dataService.GetNames()

	require.NoError(suite.T(), err)
	assert.Empty(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchUserListDatabaseError() {
	dbError := errors.New("connection timeout")
	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnError(dbError)

	result, err := suite.dataService.GetNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "db query")
	assert.Contains(suite.T(), err.Error(), dbError.Error())
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchUserListRowProcessingError() {
	invalidRows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnRows(invalidRows)

	result, err := suite.dataService.GetNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "rows scanning")
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchUserListRowIterationError() {
	iterationError := errors.New("cursor failure")
	problemRows := sqlmock.NewRows([]string{"name"}).
		AddRow("Test").
		RowError(0, iterationError)

	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnRows(problemRows)

	result, err := suite.dataService.GetNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "rows error")
	assert.Contains(suite.T(), err.Error(), iterationError.Error())
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchUserListWithClosedRows() {
	mockRows := sqlmock.NewRows([]string{"name"}).AddRow("Test").CloseError(errors.New("rows closed"))
	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnRows(mockRows)

	result, err := suite.dataService.GetNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "rows error")
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchDistinctUserListSuccess() {
	uniqueUsers := []string{"Emma", "Lucas", "Sophia"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, user := range uniqueUsers {
		mockRows.AddRow(user)
	}

	suite.mockExecutor.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(mockRows)

	result, err := suite.dataService.GetUniqueNames()

	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), uniqueUsers, result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchDistinctUserListEmptyResult() {
	emptyRows := sqlmock.NewRows([]string{"name"})
	suite.mockExecutor.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(emptyRows)

	result, err := suite.dataService.GetUniqueNames()

	require.NoError(suite.T(), err)
	assert.Empty(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchDistinctUserListDatabaseError() {
	queryError := errors.New("syntax error")
	suite.mockExecutor.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnError(queryError)

	result, err := suite.dataService.GetUniqueNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "db query")
	assert.Contains(suite.T(), err.Error(), queryError.Error())
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchDistinctUserListRowProcessingError() {
	invalidRows := sqlmock.NewRows([]string{"name"}).AddRow(nil)
	suite.mockExecutor.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(invalidRows)

	result, err := suite.dataService.GetUniqueNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "rows scanning")
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchDistinctUserListRowIterationError() {
	iterationError := errors.New("distinct cursor failure")
	problemRows := sqlmock.NewRows([]string{"name"}).
		AddRow("UniqueTest").
		RowError(0, iterationError)

	suite.mockExecutor.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(problemRows)

	result, err := suite.dataService.GetUniqueNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "rows error")
	assert.Contains(suite.T(), err.Error(), iterationError.Error())
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestFetchDistinctUserListWithClosedRows() {
	mockRows := sqlmock.NewRows([]string{"name"}).AddRow("Unique").CloseError(errors.New("distinct rows closed"))
	suite.mockExecutor.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(mockRows)

	result, err := suite.dataService.GetUniqueNames()

	require.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "rows error")
	assert.Nil(suite.T(), result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestMultipleServiceCalls() {
	firstRows := sqlmock.NewRows([]string{"name"}).AddRow("FirstUser")
	secondRows := sqlmock.NewRows([]string{"name"}).AddRow("SecondUser")

	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnRows(firstRows)
	suite.mockExecutor.ExpectQuery("SELECT DISTINCT name FROM users").
		WillReturnRows(secondRows)

	firstResult, firstErr := suite.dataService.GetNames()
	require.NoError(suite.T(), firstErr)
	assert.Equal(suite.T(), []string{"FirstUser"}, firstResult)

	secondResult, secondErr := suite.dataService.GetUniqueNames()
	require.NoError(suite.T(), secondErr)
	assert.Equal(suite.T(), []string{"SecondUser"}, secondResult)

	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestInternationalCharacterSupport() {
	internationalNames := []string{"José", "Müller", "Françoise", "Björn"}
	mockRows := sqlmock.NewRows([]string{"name"})
	for _, name := range internationalNames {
		mockRows.AddRow(name)
	}

	suite.mockExecutor.ExpectQuery("SELECT name FROM users").
		WillReturnRows(mockRows)

	result, err := suite.dataService.GetNames()

	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), internationalNames, result)
	require.NoError(suite.T(), suite.mockExecutor.ExpectationsWereMet())
}

func (suite *DatabaseOperationsTestSuite) TestDatabaseInterfaceImplementation() {
	var _ db.Database = (*sql.DB)(nil)

	mockDB := &mockDatabase{}
	service := db.New(mockDB)
	assert.NotNil(suite.T(), service)
}

type mockDatabase struct{}

func (m *mockDatabase) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}

func TestDatabaseOperationsTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseOperationsTestSuite))
}
