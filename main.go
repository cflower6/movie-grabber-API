package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-grab-movie/models"
	"io"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(CORS())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("api/movie/:title", func(c *gin.Context) {
		resp, err := http.Get("https://www.omdbapi.com/?i=tt3896198&apikey=25182b57&t=" + c.Param("title"))
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
		c.JSON(http.StatusOK, movie)
	})

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
