package main

import (
	"encoding/base64"
	"os"

	"log"

	"github.com/Strivtech/crypto-api/model"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "3324"
	}

	// initialize websites names

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(CORSMiddleware())

	router.POST("/encrypt", func(c *gin.Context) {
		var val model.Value
		c.BindJSON(&val)
		hash, _ := encrypt([]byte(val.Value), []byte(os.Getenv("SECRET_KEY")))
		res := base64.StdEncoding.EncodeToString(hash)
		c.JSON(200, gin.H{"data": res})
		return
	})
	router.POST("/decrypt", func(c *gin.Context) {
		var val model.Value
		c.BindJSON(&val)
		hash, _ := base64.StdEncoding.DecodeString(val.Value)
		res, _ := decrypt(hash, []byte(os.Getenv("SECRET_KEY")))
		c.JSON(200, gin.H{"data": string(res)})
		return
	})

	router.Run(":" + port)
}
