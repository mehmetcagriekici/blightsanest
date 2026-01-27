package search

import (
	"os"
	"fmt"
	"log"
	"time"
	"testing"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/mehmetcagriekici/blightsanest/internal/database"
)

var db *sql.DB
var queries *database.Queries

// initiate the database
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal(err)
	}

	if err := pool.Client.Ping(); err != nil {
		log.Fatal(err)
	}

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
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy(Name: "no")
	})
	if err != ni {
		log.Fatal(err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseURL := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)
	log.Println("Connecting to database on url: ", databaseURL)

	resource.Expire(120)
	pool.MaxWait = 120 * time.Second
	if err := pool.Retry(func() error {
		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
			return err
		}
		return db.Ping()
		}); err != nil {
		log.Fatal(err)
	}

	wd, _ := os.Getwd()
	log.Printf("Working dir: %s\n", wd)

	migrationsDir := "../../sql/schema"
	log.Printf("Using migrations dir: %s\n", migrationsDir)

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatal(err)
	}
}

func TestInvertedIndex(t *testing.T) {
	// initiate a new Inverted index
	invertedIndex := NewInvertedIndex()

	// test build crypto index

	// test get documents

	// test add document

	// test save document
}
