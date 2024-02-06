package api

import (
	"dummy/repository"
	"encoding/json"
	"net/http"
)

func GetModels(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.GetModels())
}
