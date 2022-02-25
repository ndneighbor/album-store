package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ndneighbor/album-store/backend/models"
)

func (h handler) GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	var books []models.Album

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
