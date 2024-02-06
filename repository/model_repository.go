package repository

import (
	"dummy/model"
)

func GetModels() ([]model.Model, error) {
	queryResult, err := Repository.DB.Query("SELECT id, attribute1 FROM model")
	if err != nil {
		return nil, err
	}

	defer queryResult.Close()

	var models []model.Model
	for queryResult.Next() {
		var m model.Model
		queryResult.Scan(&m.Id, &m.Attribute1)
		models = append(models, m)
	}

	return models, nil
}

func GetModel(id string) (model.Model, error) {
	queryResult := Repository.DB.QueryRow("SELECT id, attribute1 FROM model WHERE id = $1", id)
	var m model.Model
	err := queryResult.Scan(&m.Id, &m.Attribute1)
	return m, err
}

func DeleteModel(id string) {

}
