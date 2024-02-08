package api

import (
	"dummy/repository"
	"encoding/json"
	"log"
	"net/http"
)

type API interface {
	GetModels(w http.ResponseWriter, r *http.Request)
}

type modelAPI struct {
	repo *repository.Repository
}

func NewAPI(repo *repository.Repository) API {
	return &modelAPI{
		repo: repo,
	}
}

func (api *modelAPI) GetModels(w http.ResponseWriter, r *http.Request) {
	models, err := api.repo.GetModels()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(models)
}
