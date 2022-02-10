package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	// To be able to send tokens to the server.
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run(":8080")
}

// album type has data about the record

type album struct {
	ID     string  `json:"id"`
	Image  string  `json:"image"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Image: "https://upload.wikimedia.org/wikipedia/en/6/68/John_Coltrane_-_Blue_Train.jpg"},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99, Image: "https://upload.wikimedia.org/wikipedia/en/c/ca/Jeru_%28album%29.jpg"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Image: "https://upload.wikimedia.org/wikipedia/en/d/d8/Sarah_Vaughan.jpg"},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Add delete later
// func deleteAlbumByID(x *gin.Context) {
// 	id := c.Param("id")

// 	intVar := strconv.Atoi(id)

// 	ablums[intVar-1]
// }
