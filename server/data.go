package main

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

var albums = []album{
	{ID: "1", Title: "In The Aeroplane Over The Sea", Artist: "Neutral Milk Hotel"},
	{ID: "2", Title: "Open Your Eyes", Artist: "Goldfinger"},
	{ID: "3", Title: "The Violent Femmes (Self Title)", Artist: "The Violent Femmes"},
}
