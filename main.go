package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-grab-movie/models"
	"io"
	"log"
	"net/http"
)

var cacheMovies = make(map[string]models.SearchResults)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		api.GET("/movie", func(c *gin.Context) {
			var title = c.Query("t")
			var id = c.Query("i")

			resp, err := http.Get("http://www.omdbapi.com/?apikey=25182b57&t=" + title + "&i=" + id)
			if err != nil {
				log.Fatal(err)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var movie models.Movie
			err = json.Unmarshal(body, &movie)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(movie)
			c.JSON(http.StatusOK, movie)
		})
		api.GET("/movies", func(c *gin.Context) {
			var query = c.Query("s")
			if val, ok := cacheMovies[query]; ok {
				c.JSON(http.StatusOK, val)
			} else {
				resp, err := http.Get("http://www.omdbapi.com/?apikey=25182b57&s=" + query)
				if err != nil {
					log.Fatal(err)
				}
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				var movies models.SearchResults
				err = json.Unmarshal(body, &movies)
				if err != nil {
					log.Fatal(err)
				}
				cacheMovies[query] = movies
				log.Println(movies)
				c.JSON(http.StatusOK, movies)
			}
		})
	}

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}

}
