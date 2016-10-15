package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
)

const help = `Usage:
# Shows one def.
urban [word] 

# Limits def. to number you specify
urban [word] [limit]

# Shows all def.
urban [word] all
`

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

/* Fetches results */
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

func (r Results) print(limit int) {
	i := 0
	if r.ResultType == "exact" {
		for _, def := range r.List {
			if limit == -1 || i < limit {
				// ugly code for windows!
				color.Set(color.FgBlue)
				fmt.Printf("%d)\n", i+1)
				color.Unset()

				color.Set(color.FgGreen)
				fmt.Printf("%s\n", "Def:")
				color.Unset()
				fmt.Println(def.Definition)

				color.Set(color.FgGreen)
				fmt.Printf("%s\n", "Eg:")
				color.Unset()
				fmt.Printf("%s\n\n", def.Example)
				i++
			}
		}
	} else {
		color.Red("No definitions.\n")
		return
	}
}

/* Parses commandline arguments and prints meanings */
func ParseArg() {
	var limit int
	// argument parsing
	args, arglen := os.Args, len(os.Args)
	if arglen == 2 {
		// no limit
		limit = 1
	} else if arglen == 3 {
		// limit specified
		var err error
		// int or all
		if args[2] == "all" {
			limit = -1
		} else {
			limit, err = strconv.Atoi(args[2])
			if err != nil {
				fmt.Println("Limit should be a number.")
			}
		}
	} else {
		// for invalid args.
		fmt.Printf(help)
		return
	}
	results := GetResults(args[1])
	results.print(limit)
}

func main() {
	ParseArg()
}
