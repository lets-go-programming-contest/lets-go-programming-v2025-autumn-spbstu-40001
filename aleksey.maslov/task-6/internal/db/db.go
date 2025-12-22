package db

import (
	"database/sql"
	"fmt"
)

type Service struct {
	DB *sql.DB
}

func New(db *sql.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) GetNames() ([]string, error) {
	rows, err := s.DB.Query("SELECT name FROM users")
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("rows scanning: %w", err)
		}
		result = append(result, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return result, nil
}

func (s *Service) GetUniqueNames() ([]string, error) {
	rows, err := s.DB.Query("SELECT DISTINCT name FROM users")
	if err != nil {
		return nil, fmt.Errorf("db query: %w", err)
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("rows scanning: %w", err)
		}
		result = append(result, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return result, nil
}
