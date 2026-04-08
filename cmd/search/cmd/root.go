package cmd

// import external and internal packages
import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/search"
)

var (
	// server context
	Ctx context.Context
	// rabbitmq connection
	Conn *amqp.Connection
	// postgresql database queries
	DbQueries *database.Queries
	// semantic api client
	SemanticClient *search.Client
	// inverted index
	InvertedIndex *search.InvertedIndex
)

// prerun
var RootCmd = &cobra.Command{
	Use:   "search",
	Short: "BlightSanest search engine",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// load environment variables
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)

		}

		// create client context - shared
		Ctx = context.Background()

		// create rabbit mq connection
		conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
		if err != nil {
			log.Fatal(err)

		}

		// shared
		Conn = conn

		// create dlx for crypto
		if err := pubsub.CreateCryptoDLX(Conn); err != nil {
			log.Fatal(err)

		}

		// open the database
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)

		}

		// get database queries - shared
		DbQueries = database.New(db)

		// inverted index
		InvertedIndex = search.NewInvertedIndex()

		// semantic api client
		SemanticClient = search.NewClient(os.Getenv("SEMANTIC_API_URL"))
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
