// COPYRIGHT: https://github.com/ory/dockertest/blob/v3/examples/PostgreSQL.md

package crypto

import (
	"os"
	"log"
	"database/sql"
	"fmt"
	"testing"
	"time"
	// "context"

	"github.com/pressly/goose/v3"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

var db *sql.DB
var queries *database.Queries

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		db, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	wd, _ := os.Getwd()
	log.Printf("Working dir: %s", wd)

	migrationsDir := "../../sql/schema"
	log.Printf("Using migrations dir: %s", migrationsDir)

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("goose up: %v", err)
	}

	queries = database.New(db)

	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	// run tests
	code := m.Run()
	os.Exit(code)
}

/*
func TestCryptoDatabase(t *testing.T) {
	ctx := context.Background()
	sampleData := []MarketData{
		{
			Symbol: "BTC",
			Price:  42000.50,
		},
		{
			Symbol: "ETH",
			Price:  2300.75,
		},
	}
}
*/
