package main

import (
	//"html/template"//add html/template package 
	"net/http"
	"log"
	"fmt"
	"math/rand"
	"time"
	"regexp"

)
func elizaResponse(input string) string{
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched{
		return "why dont you tell me about your father?"
	}	

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)// or \bI'?\s*a?m\b
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?")
	}
	
	answers := []string{	
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
	}
	
	return answers[rand.Intn(len(answers))]
}	

func main() {
	// handles root page
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
	
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
	
	
}