package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

// type Intro struct {
// 	intro string    `json:"intro"`
// 	story []Stories `json:"story"`
// }

// type Stories struct {
// 	title   string    `json:"title"`
// 	story   []string  `json:"story"`
// 	options []Options `json:"options"`
// }

// type Options struct {
// 	text string `json:"text"`
// 	arc  string `json:"arc"`
// }

type StoryChapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Adventure struct {
	Intro     StoryChapter `json:"intro"`
	NewYork   StoryChapter `json:"new-york"`
	Debate    StoryChapter `json:"debate"`
	SeanKelly StoryChapter `json:"sean-kelly"`
	MarkBates StoryChapter `json:"mark-bates"`
	Denver    StoryChapter `json:"denver"`
	Home      struct {
		Title   string        `json:"title"`
		Story   []string      `json:"story"`
		Options []interface{} `json:"options"`
	} `json:"home"`
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("gopher.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteAll, _ := ioutil.ReadAll(jsonFile)

	var adven Adventure
	json.Unmarshal(byteAll, &adven)
	fmt.Println(adven.Intro.Title)
	fmt.Println(reflect.TypeOf(adven))
	tmpl := template.Must(template.ParseFiles("cyoa.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, adven)
		if err != nil {
			log.Println("Error executing template :", err)
		}
	})
	http.ListenAndServe(":80", nil)
}
