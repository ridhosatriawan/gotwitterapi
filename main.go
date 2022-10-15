package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	twitterscraper "github.com/n0madic/twitter-scraper"
)

func getUrlVideoById(id string) (*twitterscraper.Tweet, error) {
	scraper := twitterscraper.New()
	tweet, err := scraper.GetTweet(id)
	if err != nil {
		return nil, errors.New("video not found")
	}
	videos := tweet
	return videos, nil
}

func getTwitterUrlVideo(context *gin.Context) {
	id := context.Param("idTweet")

	videos, err := getUrlVideoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message ": " Video Not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, videos)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}
	router := gin.Default()
	router.GET("/twapi/:idTweet", getTwitterUrlVideo)
	router.Run("localhost:" + port)
}
