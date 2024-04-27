package farm

import "database/sql"

type Repository struct {
	db *sql.DB
}

func New(storage *sql.DB) *Repository {
	return &Repository{db: storage}
}

func (r *Repository) GetAll() *sql.Row {
	// TODO handle error
	stmt, _ := r.db.Prepare("SELECT * FROM pets")

	rows := stmt.QueryRow()

	return rows
}
