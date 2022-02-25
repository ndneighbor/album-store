package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ndneighbor/album-store/backend/pkg/models"
)

func (h handler) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.Album
	json.Unmarshal(body, &updatedBook)

	var album models.Album

	if result := h.DB.First(&album, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	album.Title = updatedBook.Title
	album.Artist = updatedBook.Artist
	album.Price = updatedBook.Price

	h.DB.Save(&album)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
