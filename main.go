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
func GetMeanings(word string, limit int) ([]string, []string) {
	var meanings []string
	var examples []string
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
	if limit == -1 {
		meanings, examples = meanings[:], examples[:]
	} else {
		limit = clamp(limit, len(meanings))
		meanings, examples = meanings[:limit], examples[:limit]
	}
	return meanings, examples
}

/* Parses commandline arguments and prints meanings */
func ParseArg() {
	var meanings []string
	var examples []string

	args, arglen := os.Args, len(os.Args)
	if arglen == 2 {
		// no limit
		meanings, examples = GetMeanings(args[1], 1)
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

		meanings, examples = GetMeanings(args[1], limit)
	} else {
		fmt.Printf(help)
		return
	}

	for i := 0; i < len(meanings); i++ {
		color.Blue("%d)\n", i+1)
		fmt.Printf("%s%s%s%s\n", color.GreenString("Def:"), meanings[i], color.GreenString("Eg:"), examples[i])
	}
}

func main() {
	ParseArg()
}
