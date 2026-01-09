package db_test

import (
	"errors"
	"testing"

	"github.com/A1exCRE/task-6/internal/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

var errTest = errors.New("test error")

const (
	selectNamesQuery  = "SELECT name FROM users"
	selectUniqueQuery = "SELECT DISTINCT name FROM users"
)

func TestNew(t *testing.T) {
	t.Parallel()

	dbConn, _, err := sqlmock.New()
	require.NoError(t, err)
	defer dbConn.Close()

	svc := db.New(dbConn)
	require.Equal(t, dbConn, svc.DB)
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rows       *sqlmock.Rows
		queryErr   error
		expected   []string
		errCheck   error
		errMessage string
	}{
		{
			rows: sqlmock.NewRows([]string{"name"}).
				AddRow("Alex").
				AddRow("Maria"),
			expected: []string{"Alex", "Maria"},
		},
		{
			rows:     sqlmock.NewRows([]string{"name"}),
			expected: nil,
		},
		{
			queryErr:   errTest,
			errCheck:   errTest,
			errMessage: "db query",
		},
		{
			rows:       sqlmock.NewRows([]string{"name"}).AddRow(nil),
			errMessage: "rows scanning",
		},
		{
			rows: sqlmock.NewRows([]string{"name"}).
				AddRow("Alex").
				AddRow("Maria").
				RowError(1, errTest),
			errCheck:   errTest,
			errMessage: "rows error",
		},
	}

	for idx, tc := range tests {
		dbConn, mock, err := sqlmock.New()
		require.NoError(t, err)

		svc := db.DBService{DB: dbConn}

		if tc.queryErr != nil {
			mock.ExpectQuery(selectNamesQuery).WillReturnError(tc.queryErr)
		} else {
			mock.ExpectQuery(selectNamesQuery).WillReturnRows(tc.rows)
		}

		got, err := svc.GetNames()

		if tc.errCheck != nil || tc.errMessage != "" {
			require.Error(t, err, "test case %d", idx)

			if tc.errCheck != nil {
				require.ErrorIs(t, err, tc.errCheck, "test case %d", idx)
			}

			if tc.errMessage != "" {
				require.ErrorContains(t, err, tc.errMessage, "test case %d", idx)
			}

			require.Nil(t, got, "test case %d", idx)
		} else {
			require.NoError(t, err, "test case %d", idx)
			require.Equal(t, tc.expected, got, "test case %d", idx)
		}

		mock.ExpectClose()
		require.NoError(t, dbConn.Close())
		require.NoError(t, mock.ExpectationsWereMet(), "test case %d", idx)
	}
}

func TestGetUniqueNames(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rows       *sqlmock.Rows
		queryErr   error
		expected   []string
		errCheck   error
		errMessage string
	}{
		{
			rows: sqlmock.NewRows([]string{"name"}).
				AddRow("Alex").
				AddRow("Maria").
				AddRow("Alex"),
			expected: []string{"Alex", "Maria", "Alex"},
		},
		{
			rows:     sqlmock.NewRows([]string{"name"}),
			expected: nil,
		},
		{
			queryErr:   errTest,
			errCheck:   errTest,
			errMessage: "db query",
		},
		{
			rows:       sqlmock.NewRows([]string{"name"}).AddRow(nil),
			errMessage: "rows scanning",
		},
		{
			rows: sqlmock.NewRows([]string{"name"}).
				AddRow("Alex").
				AddRow("Maria").
				RowError(1, errTest),
			errCheck:   errTest,
			errMessage: "rows error",
		},
	}

	for idx, tc := range tests {
		dbConn, mock, err := sqlmock.New()
		require.NoError(t, err)

		svc := db.DBService{DB: dbConn}

		if tc.queryErr != nil {
			mock.ExpectQuery(selectUniqueQuery).WillReturnError(tc.queryErr)
		} else {
			mock.ExpectQuery(selectUniqueQuery).WillReturnRows(tc.rows)
		}

		got, err := svc.GetUniqueNames()

		if tc.errCheck != nil || tc.errMessage != "" {
			require.Error(t, err, "test case %d", idx)

			if tc.errCheck != nil {
				require.ErrorIs(t, err, tc.errCheck, "test case %d", idx)
			}

			if tc.errMessage != "" {
				require.ErrorContains(t, err, tc.errMessage, "test case %d", idx)
			}

			require.Nil(t, got, "test case %d", idx)
		} else {
			require.NoError(t, err, "test case %d", idx)
			require.Equal(t, tc.expected, got, "test case %d", idx)
		}

		mock.ExpectClose()
		require.NoError(t, dbConn.Close())
		require.NoError(t, mock.ExpectationsWereMet(), "test case %d", idx)
	}
}
