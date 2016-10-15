package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Results struct {
	Tags       []string `json:"tags"`
	ResultType string   `json:"result_type"`
	List       []struct {
		Definition  string `json:"definition"`
		Permalink   string `json:"permalink"`
		ThumbsUp    int    `json:"thumbs_up"`
		Author      string `json:"author"`
		Word        string `json:"word"`
		Defid       int    `json:"defid"`
		CurrentVote string `json:"current_vote"`
		Example     string `json:"example"`
		ThumbsDown  int    `json:"thumbs_down"`
	} `json:"list"`
	Sounds []string `json:"sounds"`
}

const api_endpoint = "http://api.urbandictionary.com/v0/define?term="

func GetResults(word string) Results {
	//get json data
	res, err := http.Get(api_endpoint + word)
	if err != nil {
		log.Fatalf("Failed to fetch api.")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	//unmarshal json
	results := Results{}
	json.Unmarshal(body, &results)
	return results
}

func (r Results) print() {
	fmt.Println(r.ResultType)
}

func main() {
	r := GetResults("test")
	r.print()
}
