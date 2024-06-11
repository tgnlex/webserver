package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Handle(pattern string, handler Handler) {
	http.Handle(pattern, handler)
}

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
	c.IndentedJSON(200, gin.H{
		"Status": "Running",
	})
}
func makeTemplate(name string, html string) {
	tmpl := template.New(name)
	tmpl, err := tmpl.Parse(html)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(os.Stdout, nil)
}
