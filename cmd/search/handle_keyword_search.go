package main

import(
	"log"
	
	"github.com/mehmetcagriekici/blightsanest/internal/search"
)

func handle_keyword_search(invertedIndex *search.InvertedIndex) {
	// load the cdocuments from the cache
	if err := invertedIndex.LoadDocuments(); err != nil {
		log.Println("Before start searching, please build the inverted index with the command <create_inverted_index>")
		log.Fatal(err)
	}
}
