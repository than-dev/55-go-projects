package main

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	
}

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie