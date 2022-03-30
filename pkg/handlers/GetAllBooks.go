package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sohanpatel1221/GoAPI/pkg/mocks"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}
