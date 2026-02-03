package main

import _ "github.com/lib/pq"

import(
	"os"
	"log"
	"context"
	"database/sql"


	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"


	"github.com/mehmetcagriekici/blightsanest/internal/database"
	"github.com/mehmetcagriekici/blightsanest/internal/search"
	"github.com/mehmetcagriekici/blightsanest/internal/logs"
)

func main() {
	log.Println("Welcome to BlightSanest Search Engine...")

	// load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// get the rabbitmq and postgresql database url from .env
	rabbitURL   := os.Getenv("RABBIT_CONNECTION_STRING")
	databaseURL := os.Getenv("DB_URL")

	// create context, rabbit connection, and database queries
	ctx := context.Background()
	
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	// inverted index
	invertedIndex := search.NewInvertedIndex()

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
			if len(words) < 3 {
				log.Println("Please at least provide and asset type and a search type.")
				log.Println("search <search_type> <asset> <query...>")
				log.Println("Example: search keyword crypto find tokens with 1150 dollars.")
				continue
			}
			switch searchType := words[1]; searchType {
			case "keyword":
				// keyword search
				handle_keyword_search(invertedIndex, words[2], words[3:])
			case "semantic":
				// semantic search
				handle_semantic_search()
			}
		case "create_inverted_index":
			log.Println("Building the inverted index for the database")
			if err := invertedIndex.BuildCryptoIndex(ctx, dbQueries); err != nil {
				log.Printf("An error occured while trying to build the inverted index: %v\n", err)
				continue
			}

			log.Println("Saving the created inverted index to the local cache folder.")
			if err := invertedIndex.SaveDocuments(); err != nil {
				log.Printf("Couldn't save the created inverted index to the local machine: %v\n", err)
				continue
			}
		}
	}
}
