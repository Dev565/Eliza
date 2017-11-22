package main

import (
	//"html/template"//add html/template package 
	"net/http"
	//"log"
	"fmt"
	"math/rand"
	"time"
	"regexp"
	"path"

)
func handler(w http.ResponseWriter, r *http.Request){
		userInput := r.URL.Query().Get("user-input")
		response := elizaResponse(userInput)
		fmt.Fprintf(w, response)
}

func elizaResponse(input string) string{
	response := createReply("user-input")

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

func createReply(input string) string{
	rand.Seed(time.Now().UTC().UnixNano()) // Try changing this number!

	userInput := "inputString"
	
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched{
		return "why dont you tell me about your father?"
	}	

	re := regexp.MustCompile(`(?i)I am ([^.?!]*)[.?!]?`)// or \bI'?\s*a?m\b
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do you know you are $1?")
	}
	
	answers := []string{
		"I'm not sure what you are trying to say. Could you explain it to me?",
		"How does that make you feel?",
		 "Why do you say that?",

	}

	//returning a single string response
	return answers[rand.Intn(len(answers))]
}

// adapted from: https://github.com/ET-CS/golang-response-examples/blob/master/ajax-json.go
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    //parse request to struct
    var d Data
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func main() {

	/* console for testing regex
	fmt.Println(" Input : " + "People say I look like both my mother and father.")
		fmt.Println(" Output : " + elizaResponse("People say I look like both my mother and father"))
		fmt.Println()


		
		fmt.Println(" Input : " + "Father was a teacher.")
		fmt.Println(" Output : " + elizaResponse("Father was a teacher"))

		
		fmt.Println(" Input : " + "I was my father’s favourite.")
		fmt.Println(" Output : " + elizaResponse("I was my father’s favourite."))
		fmt.Println()


		fmt.Println(" Input : " + "Im looking forward to the weekend.")
		fmt.Println(" Output : " + elizaResponse("I’m looking forward to the weekend."))
		fmt.Println()


		fmt.Println(" Input : " + "My grandfather was French!")
		fmt.Println(" Output : " + elizaResponse("My grandfather was French!"))
		fmt.Println()

		
		fmt.Println(" Input :" + "I am happy")
		fmt.Println(" Output :" + elizaResponse("I am happy"))
		fmt.Println()


		fmt.Println(" Input :" + "I am not happy with your responses. ")
		fmt.Println(" Output :" + elizaResponse("I am not happy with your responses. "))
		fmt.Println()

				
		fmt.Println(" Input :" + " I am not sure that you understand the effect that your questions are having on me.")
		fmt.Println(" Output :" + elizaResponse("I am not sure that you understand the effect that your questions are having on me."))
		fmt.Println()

		
		fmt.Println(" Input :" + "I am supposed to just take what you’re saying at face value?")
		fmt.Println(" Output :" + elizaResponse("I am supposed to just take what you’re saying at face value? "))
		fmt.Println()
		*/
	
	// handles root page
	path := http.Dir("./static")
	fileServer := http.FileServer(path)
	
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", ajaxHandler)
	//log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
	
	
}