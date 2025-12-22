package db

import (
	"database/sql"
	"fmt"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}

type DBService struct {
	DB Database
}

func New(db Database) DBService {
	return DBService{DB: db}
}

func (service DBService) GetActiveUsers() ([]string, error) {
	query := "SELECT username FROM users WHERE active = true"
	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var users []string
	for rows.Next() {
		var user string
		if err := rows.Scan(&user); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

func (service DBService) DeactivateUser(username string) error {
	query := "UPDATE users SET active = false WHERE username = ?"
	result, err := service.DB.Exec(query, username)
	if err != nil {
		return fmt.Errorf("exec error: %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected error: %w", err)
	}

	if affected == 0 {
		return fmt.Errorf("user not found: %s", username)
	}

	return nil
}
