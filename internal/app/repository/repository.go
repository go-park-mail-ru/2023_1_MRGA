package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func New(dsn string) (*Repository, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}
