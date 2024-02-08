package repository

import (
	"database/sql"
	"dummy/model"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetModels() ([]model.Model, error) {
	queryResult, err := r.db.Query("SELECT id, attribute1 FROM model")
	if err != nil {
		return nil, err
	}

	defer queryResult.Close()

	var models []model.Model
	for queryResult.Next() {
		var m model.Model
		err = queryResult.Scan(&m.Id, &m.Attribute1)
		if err != nil {
			return nil, err
		}
		models = append(models, m)
	}

	return models, nil
}

func (r *Repository) GetModel(id string) (model.Model, error) {
	queryResult := r.db.QueryRow("SELECT id, attribute1 FROM model WHERE id = $1", id)
	var m model.Model
	err := queryResult.Scan(&m.Id, &m.Attribute1)
	return m, err
}

func (r *Repository) DeleteModel(id string) {
	panic("not implemented yet")
}
