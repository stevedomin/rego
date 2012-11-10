package main

import (
	"encoding/json"
	// "fmt"
	"html/template"
	"log"
	"net/http"
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
	regexString := req.FormValue("regex")
	testString := req.FormValue("testString")

	m := &MatchResult{}

	regex, _ := regexp.Compile(regexString)
	matches := regex.FindStringSubmatch(testString)

	if len(matches) > 0 {
		m.Match = matches[0]
		m.Groups = matches[1:]
	}

	m.GroupsName = regex.SubexpNames()[1:]

	enc := json.NewEncoder(rw)
	enc.Encode(m)

	// fmt.Fprintf(rw, "%t", match)

}

func main() {
	// Main handler (index.html)
	http.HandleFunc("/", handler)
	// Regex testing service
	http.HandleFunc("/test_regex/", regExpHandler)
	// Static file serving
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Launch server
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
