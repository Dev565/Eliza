package main

import (
	//"html/template"//add html/template package 
	"net/http"

)
	
type myMsg struct {
	Input string
	Output string
	Previous string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	//serve the homepage.html file
	http.ServeFile(w, r, "eliza.html")
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	response := "hello"
	message := r.FormValue("chat")
	previous := input 

}//chatHandler

func main() {
	// handles root page
	http.HandleFunc("/", requestHandler)

	//handle /chat page
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":8080", nil)
}