package main

import "net/http"

type res = http.ResponseWriter
type req = http.Request

type fetchResult struct {
	Message string
	Error   error
}

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

type Handler interface {
	ServeHTTP(res, *req)
}
type HandlerFunc func(res, *req)