package repository

import (
	"dummy/mock"
	"dummy/model"
)

func GetModels() []model.Model {
	models := mock.MockModels()
	return models
}

func GetModel(id string) model.Model {
	m := mock.MockModel(id, "ASDASD")
	return m
}
