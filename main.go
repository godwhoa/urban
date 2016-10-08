package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
	"log"
	"os"
	"strconv"
)

const help = `Usage:
# Shows one def.
urban [word] 

# Limits def. to number you specify
urban [word] [limit]

# Shows all def.
urban [word] all
`

/* If statement spaghetti from here on down. */

/* Prevent going above highest limit */
func clamp(n, high int) int {
	if n > high {
		return high
	} else {
		return n
	}
}

/* Returns meanings and examples */
func GetMeanings(word string, limit int) ([]string, []string, bool) {
	var meanings []string
	var examples []string
	var hasresults = true
	doc, err := goquery.NewDocument(`http://www.urbandictionary.com/define.php?term=` + word)
	if err != nil {
		log.Fatalf("Failed to fetch site.\n")
	}

	doc.Find(".meaning").Each(func(i int, s *goquery.Selection) {
		meanings = append(meanings, s.Text())
	})

	doc.Find(".example").Each(func(i int, s *goquery.Selection) {
		examples = append(examples, s.Text())
	})
	// we use -1 to indicate we want all results
	if limit == -1 {
		meanings, examples = meanings[:], examples[:]
	} else if len(meanings) == len(examples) {
		limit = clamp(limit, len(meanings))
		meanings, examples = meanings[:limit], examples[:limit]
	} else {
		hasresults = false
	}
	return meanings, examples, hasresults
}

/* Parses commandline arguments and prints meanings */
func ParseArg() {
	var meanings []string
	var examples []string
	var hasresults bool

	// argument parsing
	args, arglen := os.Args, len(os.Args)
	if arglen == 2 {
		// no limit
		meanings, examples, hasresults = GetMeanings(args[1], 1)
	} else if arglen == 3 {
		// limit specified
		var limit int
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

		meanings, examples, hasresults = GetMeanings(args[1], limit)
	} else {
		// for invalid args.
		fmt.Printf(help)
		return
	}

	// only loop when it has results
	if hasresults {
		for i := 0; i < len(meanings); i++ {
			color.Blue("%d)\n", i+1)
			fmt.Printf("%s%s%s%s\n", color.GreenString("Def:"), meanings[i], color.GreenString("Eg:"), examples[i])
		}
	} else {
		color.Red("No definitions.\n")
	}
}

func main() {
	ParseArg()
}
