package ui

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/pivotal-cf/guess/guess"
	"math/rand"
)

func Webserver () {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/guess", func(c *gin.Context){
		t := guess.GenerateTarget(rand.Intn)
		guessedNumber, _ := strconv.Atoi(c.Query("value"))

		result := guess.IsGuessCorrect(guessedNumber, t)
		c.JSON(200, gin.H{
			"your-guess": guessedNumber,
			"result":result,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}