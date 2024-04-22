package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Intro struct {
	intro string    `json:"intro"`
	story []Stories `json:"story"`
}

type Stories struct {
	title   string    `json:"title"`
	story   []string  `json:"story"`
	options []Options `json:"options"`
}

type Options struct {
	text string `json:"text"`
	arc  string `json:"arc"`
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

	var intro Intro
	json.Unmarshal(byteAll, &intro)
	fmt.Println(intro)
}
