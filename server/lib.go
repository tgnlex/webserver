package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getTime() {
	print(timestamp)
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func fetch(data string) fetchResult {
	sleep(time.Second * 4)
	return fetchResult{data, nil}
}
func startup() {
	var start = fetch(STARTUP)
	print(start.Message)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
