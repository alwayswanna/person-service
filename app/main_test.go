package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"net/http"
	"person-service/model"
	"person-service/utils"
	"strings"
	"testing"
	"time"
)

var postgresContainer, ctx = initPostgresContainerAndContext()

func Test_PersonService(t *testing.T) {

	/* init server */
	server := setupHttpServer(configuration, router)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("Http-s start failed, ", utils.Err(err))
		}
	}()

	client := &http.Client{}

	t.Run("must return 200 when create person", func(t *testing.T) {
		resp, err := http.Post(
			"http://localhost:9902/api/v1/person/create",
			"application/json",
			bytes.NewBuffer([]byte(
				`{
					"firstName": "Алексей",
    				"lastName": "Сидоров",
    				"age": 18
				}`,
			)),
		)

		result := parseResponse(err, resp, t)

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "Алексей", result.FirstName)
		assert.Equal(t, "Сидоров", result.LastName)
	})

	t.Run("must return 200 when update person", func(t *testing.T) {
		resp, err := http.Post(
			"http://localhost:9902/api/v1/person/create",
			"application/json",
			bytes.NewBuffer([]byte(
				`{
					"firstName": "Алексей",
    				"lastName": "Сидоров",
    				"age": 18
				}`,
			)),
		)

		responseBytes := parseResponseBytes(err, t, resp)
		responseString := strings.Replace(string(responseBytes), "Сидоров", "Петров", 1)

		req, err := http.NewRequest(http.MethodPut, "http://localhost:9902/api/v1/person/update", strings.NewReader(responseString))
		req.Header.Set("Content-Type", "application/json")

		resp, err = client.Do(req)

		result := parseResponse(err, resp, t)

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "Алексей", result.FirstName)
		assert.Equal(t, "Петров", result.LastName)
	})

	t.Run("must return 200 when find person", func(t *testing.T) {
		resp, err := http.Post(
			"http://localhost:9902/api/v1/person/create",
			"application/json",
			bytes.NewBuffer([]byte(
				`{
					"firstName": "Петр",
    				"lastName": "Сидоров",
    				"age": 21
				}`,
			)),
		)

		result := parseResponse(err, resp, t)

		resp, err = http.Get(fmt.Sprintf("http://localhost:9902/api/v1/person/get/%s", result.Id.String()))

		result = parseResponse(err, resp, t)

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "Петр", result.FirstName)
		assert.Equal(t, "Сидоров", result.LastName)
	})

	t.Run("must return 200 when delete person", func(t *testing.T) {
		resp, err := http.Post(
			"http://localhost:9902/api/v1/person/create",
			"application/json",
			bytes.NewBuffer([]byte(
				`{
					"firstName": "Алексей",
    				"lastName": "Сидоров",
    				"age": 18
				}`,
			)),
		)

		parsedResponse := parseResponse(err, resp, t)
		req, err := http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("http://localhost:9902/api/v1/person/delete/%s", parsedResponse.Id.String()),
			nil,
		)

		resp, err = client.Do(req)

		result := parseResponseBytes(err, t, resp)

		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(
			t,
			fmt.Sprintf("{\"message\":\"Person with id: %s was successfully deleted\"}\n", parsedResponse.Id.String()),
			string(result),
		)
	})

	/* close postgres postgresContainer */
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
}

// initPostgresContainerAndContext method create new postgresContainer with postgres and initialize context.
func initPostgresContainerAndContext() (*postgres.PostgresContainer, context.Context) {
	/* init data base */
	var backgroundContext = context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "docker.io/postgres:15.2-alpine",
		ExposedPorts: []string{"5432:5432"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
	}

	postgresContainer, err := postgres.RunContainer(
		backgroundContext,
		testcontainers.CustomizeRequest(testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          false,
			ProviderType:     0,
			Logger:           nil,
			Reuse:            false,
		}),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		panic(err)
	}

	return postgresContainer, backgroundContext
}

func parseResponse(err error, response *http.Response, t *testing.T) *model.PersonResponse {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	resBody := parseResponseBytes(err, t, response)

	var res *model.PersonResponse

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		t.Fatalf("Error while parse response: %v", err)
	}

	return res
}

func parseResponseBytes(err error, t *testing.T, response *http.Response) []byte {
	if err != nil {
		t.Fatalf("Error while parse response: %v", err)
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Error while parse response: %v", err)
	}

	return resBody
}
