package main

import (
	//"html/template"//add html/template package 
	"net/http"
	"log"

)

func main() {
	// handles root page
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
	
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
	
	
}