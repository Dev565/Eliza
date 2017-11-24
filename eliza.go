package main
//  adapted from: https://github.com/data-representation/eliza/blob/master/eliza.go
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"net/http"
)

type Replacer struct {
	original *regexp.Regexp
	replacements []string
}

/*

func handler(w http.ResponseWriter, r *http.Request){
		userInput := r.URL.Query().Get("user-input")
		response := elizaResponse(userInput)
		fmt.Fprintf(w, response)
}
*/

func ReadFromFile(path string) []Replacer{
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		
		var replacers []Replacer
		
		for scanner, readoriginal := bufio.NewScanner(file), false; scanner.Scan();{
			
			switch line := scanner.Text();{

				case strings.HasPrefix(line, "#"):
				
				case len(line) == 0:
					readoriginal = false
					
				case readoriginal == false:
					replacers = append(replacers, Replacer{original: regexp.MustCompile(line)})
					readoriginal = true
				
				default:
					replacers[len(replacers)-1].replacements = append(replacers[len(replacers)-1].replacements, line)
			}
		}
	return replacers 
}

type Eliza struct {
		responses		[]Replacer
		substitutions 	[]Replacer
}

func ElizaFromFiles(responsePath string, substitutionPath string) Eliza {
	eliza := Eliza{}
	
	eliza.responses = ReadFromFile(responsePath)
	eliza.substitutions = ReadFromFile(substitutionPath)
	
	return eliza
}

func (me *Eliza) RespondTo(input string) string{

	for _, response := range me.responses {
		if matches := response.original.FindStringSubmatch(input); matches != nil {
	
			output := response.replacements[rand.Intn(len(response.replacements))]
		
			bounderies := regexp.MustCompile(`[\s,.?!]+`)
		
			for m, match := range matches[1:] {
				tokens := bounderies.Split(match, -1)
				for t, token := range tokens {
					for _, substitution := range me.substitutions {
						if substitution.original.MatchString(token) {
							tokens[t] = substitution.replacements[rand.Intn(len(substitution.replacements))]
							break
						}
					}
				}
		output = strings.Replace(output, "$"+strconv.Itoa(m+1), strings.Join(tokens, " "), -1)
	}
	return output
		}
	}
	return "I'm woldn't know too much about that"
}


	/*
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
	
func createReply(path string) string{
	rand.Seed(time.Now().UTC().UnixNano()) // Try changing this number!

	//userInput := inputString
	
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
*/

func inputHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("userInput")
print("input isin inputHandler "+input)
	eliza := ElizaFromFiles("txt/responses.txt", "txt/substitutions.txt")

	fmt.Println("Eliza: Hello, how are you today?")
output := eliza.RespondTo(input)
	fmt.Fprintf(w, output)
}

//  adapted from: https://github.com/data-representation/eliza/blob/master/eliza.go
func main() {
	// Seed the rand package with the current time.
	rand.Seed(time.Now().UnixNano())

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/chat", inputHandler)

	http.ListenAndServe(":8080", nil)

	// Create a new instance of Eliza.
//	eliza := ElizaFromFiles("txt/responses.txt", "txt/substitutions.txt")

	// Print a greeting to the user.
//	fmt.Println("Eliza: Hello, how are you today?")
	// Read from the user.
/*	scanner := bufio.NewScanner(os.Stdin)
	for fmt.Print("You: "); scanner.Scan(); fmt.Print("You: ") {
		// Print Eliza's response.
		fmt.Println("Eliza:", eliza.RespondTo(scanner.Text()))
		// If the user typed "quit" then exit. Eliza has a chance to respond first.
		if quit, _ := regexp.MatchString("(?i)^quit$", scanner.Text()); quit {
			break
		}
	}*/
}
/*
// adapted from: https://github.com/ET-CS/golang-response-examples/blob/master/ajax-json.go
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    //parse request to struct
    var d Data
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
*/
