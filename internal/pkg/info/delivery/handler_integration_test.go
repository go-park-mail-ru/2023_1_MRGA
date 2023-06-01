package delivery

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func openDB() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository:   "mdillon/postgis",
		Tag:          "latest",
		Env:          []string{"POSTGRES_PASSWORD=mysecretpassword"},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: "5433"},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	databaseConnStr := fmt.Sprintf("host=localhost port=5433 user=postgres dbname=postgres password=mysecretpassword sslmode=disable")

	log.Println("Connecting to database on url: ", databaseConnStr)

	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {

		db, err = gorm.Open(postgres.Open(databaseConnStr), &gorm.Config{})

		//
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	print("ok")

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

}

func TestHandler_GetHashtags(t *testing.T) {

	log.Println("Initialize test database...")
	openDB()
	log.Println("Create new Iris app...")

	print("app ok")
}
