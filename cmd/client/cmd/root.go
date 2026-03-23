package cmd

// import external and internal packages
import (
        "os"
	"log"
	"time"
	"context"
	"strconv"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"

        "github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/pubsub"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

var (
	// client context
        Ctx         context.Context
	// rabbitmq connection
	Conn        *amqp.Connection
	// client crypto state
	CryptoState *crypto.CryptoState
	// client crypto cache
	CryptoCache *crypto.CryptoCache
	// subscription manager for crypto
	SubManager  *pubsub.SubscriptionManager
	// postgresql database queries
	DbQueries   *database.Queries
)

// shared state and persisten prerun
var RootCmd = &cobra.Command{
	Use:   "client",
	Short: "BlightSanest crypto client",
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

		// create the cache interval
		interval, err := strconv.ParseFloat(os.Getenv("CACHE_INTERVAL"), 64)
		if err != nil {
			log.Fatal(err)

		}

		// create shared crypto state, sub manager and cache
		CryptoState = crypto.CreateCryptoState()
		SubManager  = pubsub.NewSubscriptionManager()
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
