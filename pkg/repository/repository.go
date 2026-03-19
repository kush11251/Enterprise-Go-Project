package repository

import (
	"database/sql"
)

// Repository handles data storage
func NewRepository() *Repository {
	return &Repository{db: sql.Open("sqlite3", "./example.db")}
}

type Repository struct {
	db *sql.DB
}

// SaveModel saves a model to the database
func (r *Repository) SaveModel(model model.Model) {
	// Mock saving model to database
	stmt, _ := r.db.Prepare("INSERT INTO models (name, age) VALUES (?, ?)")
	_, _ = stmt.Exec(model.Name, model.Age)
}