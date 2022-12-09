package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

var db *sql.DB

func TestResponseWithLove(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to docker")

	// resource, err := pool.Run("docker-gs-ping", "latest", []string{})
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name:         "docker-gs-ping",
		Repository:   "docker-gs-ping",
		ExposedPorts: []string{"8888", "8888"},
		Tag:          "latest",
	})
	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "vailed to remove container")
	})

	var resp *http.Response

	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetHostPort("8888/tcp"), "/"))
		if err != nil {
			t.Log("container not ready, waiting....")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	// Finally, test the business requirement
	require.Contains(t, string(body), "<3", "does'n respond with love")
}

// func TestMain(m *testing.M) {
// 	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}

// 	// pulls an image, creates a container based on it and runs it
// 	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
// 		Repository: "postgres",
// 		Tag:        "11",
// 		Env: []string{
// 			"POSTGRES_PASSWORD=secret",
// 			"POSTGRES_USER=user_name",
// 			"POSTGRES_DB=dbname",
// 			"listen_addresses = '*'",
// 		},
// 	}, func(config *docker.HostConfig) {
// 		// set AutoRemove to true so that stopped container goes away by itself
// 		config.AutoRemove = true
// 		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
// 	})
// 	if err != nil {
// 		log.Fatalf("Could not start resource: %s", err)
// 	}

// 	hostAndPort := resource.GetHostPort("5432/tcp")
// 	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

// 	log.Println("Connecting to database on url: ", databaseUrl)

// 	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

// 	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
// 	pool.MaxWait = 120 * time.Second
// 	if err = pool.Retry(func() error {
// 		db, err = sql.Open("postgres", databaseUrl)
// 		if err != nil {
// 			return err
// 		}
// 		return db.Ping()
// 	}); err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}
// 	//Run tests
// 	code := m.Run()

// 	// You can't defer this because os.Exit doesn't care for defer
// 	if err := pool.Purge(resource); err != nil {
// 		log.Fatalf("Could not purge resource: %s", err)
// 	}

// 	os.Exit(code)
// }

// func TestConnection(t *testing.T) {
// 	require.NotEmpty(t, db)
// }
