package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type MatchResultResponse struct {
	Matches    [][]string `json:"matches"`
	GroupsName []string   `json:"groupsName"`
}

func handler(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(rw, nil)
}

func regExpHandler(rw http.ResponseWriter, req *http.Request) {
	var matches [][]string

	req.ParseForm()
	regexpString := req.FormValue("regexp")
	testString := req.FormValue("testString")
	findAll, _ := strconv.ParseBool(req.FormValue("findAll"))

	log.Printf("Regexp : %s", regexpString)
	log.Printf("Test string : %s", testString)
	log.Printf("Find all : %t", findAll)

	m := &MatchResultResponse{}

	r, err := regexp.Compile(regexpString)
	if err != nil {
		log.Printf("Invalid RegExp : %s \n", regexpString)
		rw.WriteHeader(500)
		fmt.Fprintf(rw, "Invalid RegExp : %s", regexpString)
		return
	}

	if findAll {
		matches = r.FindAllStringSubmatch(testString, -1)
	} else {
		matches = [][]string{r.FindStringSubmatch(testString)}
	}

	log.Println(matches)

	if len(matches) > 0 {
		m.Matches = matches
		m.GroupsName = r.SubexpNames()[1:]
	}

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
