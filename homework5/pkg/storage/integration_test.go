package storage_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/alextonkonogov/gb_go_postgres/homework5/pkg/config"
	"github.com/alextonkonogov/gb_go_postgres/homework5/pkg/storage"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var db *sql.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13.1",
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

	hostAndPort := resource.GetHostPort("5435/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second

	ctx := context.Background()
	dbpool, err := storage.InitDBConn(ctx, &config.AppConfig{1, 1, true})
	if err != nil {
		return
	}
	defer dbpool.Close()

	log.Println("Creating tables and inserting data to database.")
	err = storage.InitTables(ctx, dbpool)
	if err != nil {

	}

	requests := []string{
		"Кинг", "Достоевский", "Лондон", "Рэнд",
	}

	for _, request := range requests {
		authors, err := storage.GetAuthorBySurname(ctx, dbpool, request)
		if err != nil {
			log.Fatalf("error while finding author: %s", err)
		}

		if len(authors) == 0 {
			log.Fatalf("Could not find author by your request: %s", err)
			return
		}

		for _, a := range authors {
			books, err := storage.GetBooksByAuthorId(ctx, dbpool, a.Id)
			if err != nil {
				log.Fatalf("Could not find books: %s", err)
			}

			fmt.Printf("По запросу \"%s\" найдены книги:\n", request)
			for _, v := range books {
				fmt.Printf("- %s\n", v.Title)
			}
		}
	}

	//Run tests
	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
