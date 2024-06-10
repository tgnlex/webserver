package main 


type fetchResult struct {
	Message string
	Error   error
}

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}