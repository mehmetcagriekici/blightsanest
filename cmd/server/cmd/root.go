package cmd

// import external and internal packages
import (
	"context"
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/cobra"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

var (
	// server context
	Ctx context.Context
	// rabbitmq connection
	Conn *amqp.Connection
	// server crypto cache
	CryptoCache *crypto.CryptoCache
	// postgresql database queries
	DbQueries *database.Queries
	// crypto API Key
	CryptoAPIKey string
)

// prerun
var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "BlightSanest crypto server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// load environment variables
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)

		}

		// CoinGecko API Key
		CryptoAPIKey = os.Getenv("COIN_GECKO_KEY")

		// create client context - shared
		Ctx = context.Background()

		// create rabbit mq connection
		conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
		if err != nil {
			log.Fatal(err)

		}

		// shared
		Conn = conn

		// create the cache interval
		interval, err := strconv.ParseFloat(os.Getenv("CACHE_INTERVAL"), 64)
		if err != nil {
			log.Fatal(err)

		}

		// create shared crypto cache
		CryptoCache = crypto.CreateCryptoCache(time.Duration(interval) * time.Hour)

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
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
