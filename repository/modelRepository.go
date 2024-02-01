package repository

import "dummy/model"

func GetModels() []model.Model {
	var models []model.Model
	models = append(models, GetModel("id1"))
	return models
}

func GetModel(id string) model.Model {
	var m model.Model
	m.SetId(id)
	m.SetAttribute1("some string")
	return m
}
