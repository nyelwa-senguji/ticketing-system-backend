package db

import "database/sql"

// Repository provides all functions to execute db queries and transcations
type Repository struct {
	*Queries
	db *sql.DB
}

// New Repository creates a new repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db:      db,
		Queries: New(db),
	}
}
