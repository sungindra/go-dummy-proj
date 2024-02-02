package api

import (
	"dummy/mock"
	"encoding/json"
	"net/http"
)

func GetModels(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mock.MockModels())
}
