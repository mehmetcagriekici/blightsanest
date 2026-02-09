package main

import(
	"fmt"
	"log"
	"strings"
	"strconv"
	
	"github.com/mehmetcagriekici/blightsanest/internal/search"
)

func handle_semantic_search(client *search.Client, limit string, queryArray []string) {
	n, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatal(err)
	}

	q := strings.Join(queryArray, " ")
	res, err := client.Search(q, n)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range res.Results {
		fmt.Printf("Score: %f, Doc ID: %s\n", r.Score, r.Document.ID)
	}
}
