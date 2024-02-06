package api

import (
	"dummy/repository"
	"encoding/json"
	"log"
	"net/http"
)

func GetModels(w http.ResponseWriter, r *http.Request) {
	models, err := repository.GetModels()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(models)
}
