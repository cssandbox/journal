package journal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// LambdaHandler Creates a Lambda Handler
func LambdaHandler(j Journal) interface{} {
	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch {
		case req.HTTPMethod == "POST":

			if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
				return clientError(http.StatusNotAcceptable), nil
			}
			en := new(Entry)
			err := json.Unmarshal([]byte(req.Body), en)
			if err != nil {
				return clientError(http.StatusUnprocessableEntity), err
			}

			return handlePost(en, j)

		case req.HTTPMethod == "GET":
			entries, err := j.read(req.QueryStringParameters["uuid"])
			if err != nil {
				return serverError(err), err
			}
			return handleGet(entries)

		case req.HTTPMethod == "PUT":
			return j.update(req)

		default:
			return clientError(http.StatusMethodNotAllowed), nil
		}
	}
}

func handlePost(en *Entry, j Journal) (events.APIGatewayProxyResponse, error) {
	if en.Title == "" || en.Text == "" {
		return clientError(http.StatusBadRequest), fmt.Errorf("Title or text is empty")
	}

	uuid, err := j.create(en)
	if err != nil {
		return serverError(err), err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers: map[string]string{
			"Location":                    fmt.Sprintf("/entries?uuid=%s", uuid),
			"Access-Control-Allow-Origin": "https://www.chaiappsPost.com",
		},
	}, nil

}

func handleGet(entries interface{}) (events.APIGatewayProxyResponse, error) {
	if entries == nil {
		return clientError(http.StatusNotFound), nil
	}
	js, err := json.Marshal(entries)
	if err != nil {
		return serverError(err), err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
		Headers:    map[string]string{"Access-Control-Allow-Origin": "https://www.chaiapps.com"},
	}, nil

}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) events.APIGatewayProxyResponse {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}
}

// Similarly add a helper for send responses relating to client errors.
func clientError(status int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}
}
