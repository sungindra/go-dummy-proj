package mock

import "dummy/model"

func MockModel(id string, attribute1 string) model.Model {
	var m model.Model
	m.Id = id
	m.Attribute1 = attribute1
	return m
}

func MockModels() []model.Model {
	var models []model.Model
	models = append(models, MockModel("id1", "Some String"), MockModel("id2", "Some Other String"))
	return models
}
