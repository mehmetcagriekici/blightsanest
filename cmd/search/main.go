package main

import(
	"log"

	"github.com/mehmetcagriekici/blightsanest/internal/search"
	"github.com/mehmetcagriekici/blightsanest/internal/logs"
)

func main() {
	log.Println("Welcome to BlightSanest Search Engine...")

	// REPL
	for {
		words := logs.GetInput()
		if len(words) == 0 {
			log.Println("To use the search engine:")
			search.PrintSearchHelp()
			continue
		}

		switch cmd := words[0]; cmd {
		case "search":
			if len(words) < 2 {
				log.Println("Please provide a search type.")
				switch searchType := words[1]; searchType {
				case "keyword":
				}
			}
		}
	}
}
