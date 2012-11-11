package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type MatchResult struct {
	Match      string   `json:"match"`
	GroupsName []string `json:"groupsName"`
	Groups     []string `json:"groups"`
}

func handler(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(rw, nil)
}

func regExpHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	regexpString := req.FormValue("regexp")
	testString := req.FormValue("testString")

	log.Printf("Regexp : %s", regexpString)
	log.Printf("Test string : %s", testString)

	m := &MatchResult{}

	r, _ := regexp.Compile(regexpString)
	matches := r.FindStringSubmatch(testString)

	if len(matches) > 0 {
		m.Match = matches[0]
		m.Groups = matches[1:]
	}

	m.GroupsName = r.SubexpNames()[1:]

	enc := json.NewEncoder(rw)
	enc.Encode(m)

}

func main() {
	// Main handler (index.html)
	http.HandleFunc("/", handler)
	// Regex testing service
	http.HandleFunc("/test_regexp/", regExpHandler)
	// Static file serving
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Launch server
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
