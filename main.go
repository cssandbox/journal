package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
)

// var isbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type entry struct {
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return create(req)
	case "GET":
		return read(req)
	case "PUT":
		return update(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func read(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the `isbn` query string parameter from the request and
	// validate it.
	uuid := req.QueryStringParameters["uuid"]
	// if !isbnRegexp.MatchString(isbn) {
	// 	return clientError(http.StatusBadRequest)
	// }

	var js []byte
	if uuid == "" {
		entries, err := getItems()
		if err != nil {
			return serverError(err)
		}
		if entries == nil {
			return clientError(http.StatusNotFound)
		}
		js, err = json.Marshal(entries)
		if err != nil {
			return serverError(err)
		}
	} else {
		// Fetch the journal entry record from the database based on the uuid value.
		en, err := getItem(uuid)
		if err != nil {
			return serverError(err)
		}
		if en == nil {
			return clientError(http.StatusNotFound)
		}

		// The APIGatewayProxyResponse.Body field needs to be a string, so
		// we marshal the journal entry record into JSON.
		js, err = json.Marshal(en)
		if err != nil {
			return serverError(err)
		}

	}

	// Return a response with a 200 OK status and the JSON book record
	// as the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	en := new(entry)
	err := json.Unmarshal([]byte(req.Body), en)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if en.Title == "" || en.Body == "" {
		return clientError(http.StatusBadRequest)
	}

	en.UUID = uuid.New().String()

	err = putItem(en)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/entries?uuid=%s", en.UUID)},
	}, nil
}

func update(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	uuid := req.QueryStringParameters["uuid"]
	if uuid == "" {
		return clientError(http.StatusBadRequest)
	}

	// Fetch the journal entry record from the database based on the uuid value.
	en, err := getItem(uuid)
	if err != nil {
		return serverError(err)
	}
	if en == nil {
		return clientError(http.StatusNotFound)
	}

	updateEn := new(entry)
	err = json.Unmarshal([]byte(req.Body), updateEn)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if updateEn.Title != "" {
		en.Title = updateEn.Title
	}

	if updateEn.Body != "" {
		en.Body = updateEn.Body
	}

	err = putItem(en)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/entries?uuid=%s", en.UUID)},
	}, nil
}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	// fmt.Println(uuid.New())
	lambda.Start(router)
}
