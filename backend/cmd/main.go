package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ndneighbor/album-store/backend/pkg/db"
	"github.com/ndneighbor/album-store/backend/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/albums", h.GetAllAlbums).Methods(http.MethodGet)
	router.HandleFunc("/albums/{id}", h.GetAlbum).Methods(http.MethodGet)
	router.HandleFunc("/albums", h.AddAlbum).Methods(http.MethodPost)
	router.HandleFunc("/albums/{id}", h.UpdateAlbum).Methods(http.MethodPut)
	router.HandleFunc("/albums/{id}", h.DeleteAlbum).Methods(http.MethodDelete)

	log.Println("Your API is live!")
	http.ListenAndServe(":4000", router)
}
