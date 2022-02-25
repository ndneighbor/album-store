package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ndneighbor/album-store/backend/pkg/models"
)

func (h handler) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find the album by Id

	var album models.Album

	if result := h.DB.First(&album, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that album
	h.DB.Delete(&album)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
