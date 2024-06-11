package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var env = os.Getenv
var timestamp = time.Now
var sleep = time.Sleep
var print = fmt.Println
var httpStat = http.StatusCreated
var ok = http.StatusOK


