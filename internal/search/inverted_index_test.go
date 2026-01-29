package search

import (
	"os"
	"fmt"
	"log"
	"time"
	"context"
	"testing"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/mehmetcagriekici/blightsanest/internal/crypto"
	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

var db *sql.DB
var queries *database.Queries

// initiate the database
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

func TestInvertedIndex(t *testing.T) {
	ctx := context.Background()
	
	// initiate a new Inverted index
	invertedIndex := NewInvertedIndex()

	// sample data
	sampleData := []crypto.MarketData{
		{
			Symbol: "BTC",
			CurrentPrice: 123.456,
		},
		{
			Symbol: "ETH",
			CurrentPrice: 456.789,
		},
	}
	sampleKey := "sample_key"
	sampleQuery := "find sample key"
	// upload sample data to the database before tests
	if err := crypto.CreateCryptoRow(ctx, queries, sampleData, sampleKey); err != nil {
		t.Errorf("While trying to upload sample data to the database for inverted index testing, an unexpected error happened: %v\n", err)
	}

	// test build crypto index
	if err := invertedIndex.BuildCryptoIndex(ctx, queries); err != nil {
		t.Errorf("Unexpected error while trying to build inverted index for crypto: %v\n", err)
	}
	
	// test get documents
	docs := invertedIndex.GetDocuments(sampleQuery)
	if docs == nil {
		t.Errorf("Get documents unexpectedly returns nil.")
	}

	// test save document
	nIdx, nDoc, nTf, err := invertedIndex.SaveDocuments()
	if err != nil {
		t.Errorf("Unexpected error while trying to save the inverted index into local files: %v\n", err)
	}

	if nIdx == 0 {
		t.Errorf("Zero amount of bytes received for inverted index index")
	}

	if nDoc == 0 {
		t.Errorf("Zero amount of bytes received for inverted index docmap")
	}

	if nTf == 0 {
		t.Errorf("Zero amount of bytes received for inverted index term frequencies")
	}

	// test load documents
	if err := invertedIndex.LoadDocuments(); err != nil {
		t.Errorf("Unexpected error while trying to load the inverted index from the local cache: %v\n", err)
	}
}
