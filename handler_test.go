package journal

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func Setup() Router {
	var mockSore Store
	return InitializeHandler(mockSore)
}

func TestCreateFailsForWrongContentType(t *testing.T) {
	router := Setup()

	req := events.APIGatewayProxyRequest{}
	req.HTTPMethod = "POST"
	req.Headers = make(map[string]string)
	req.Headers["Content-Type"] = "text"

	resp, _ := router(req)

	assert.Equal(t, http.StatusNotAcceptable, resp.StatusCode, "Expected a 406")
}

func TestCreateFailsForNonJsonBody(t *testing.T) {
	router := Setup()

	req := events.APIGatewayProxyRequest{}
	req.HTTPMethod = "POST"
	req.Headers = make(map[string]string)
	req.Headers["Content-Type"] = "application/json"
	req.Body = "Not a Json Text"

	resp, _ := router(req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, "Expected a 422")
}

func TestCreateMissingFieldsCauseBadRequest(t *testing.T) {
	router := Setup()

	req := events.APIGatewayProxyRequest{}
	req.HTTPMethod = "POST"
	req.Headers = make(map[string]string)
	req.Headers["Content-Type"] = "application/json"
	req.Body = "{\"title\":\"Test Title\"}"

	resp, _ := router(req)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected a 400")
}
