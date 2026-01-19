package main

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
	if err = godotenv.Load(); err != nil {
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
		case "create_inverted_index":
			log.Prinln("")
		}
	}
}
