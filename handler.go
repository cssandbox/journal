package journal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
)

// var isbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

// Journal struct
type Journal struct {
	Store Store
}

func (j *Journal) read(uuid string) (e interface{}, err error) {
	if uuid == "" {
		return j.Store.GetItems()
	}
	return j.Store.GetItem(uuid)
}

func (j *Journal) create(en *Entry) (string, error) {
	en.UUID = uuid.New().String()
	return en.UUID, j.Store.PutItem(en)
}

func (j *Journal) update(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable), nil
	}

	uuid := req.QueryStringParameters["uuid"]
	if uuid == "" {
		return clientError(http.StatusBadRequest), fmt.Errorf("uuid is empty")
	}

	// Fetch the journal entry record from the database based on the uuid value.
	en, err := j.Store.GetItem(uuid)
	if err != nil {
		return serverError(err), err
	}
	if en == nil {
		return clientError(http.StatusNotFound), fmt.Errorf("No item found")
	}

	updateEn := new(Entry)
	err = json.Unmarshal([]byte(req.Body), updateEn)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity), err
	}

	updateEn.UUID = en.UUID

	err = j.Store.PutItem(updateEn)
	if err != nil {
		return serverError(err), err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "https://www.chaiapps.com"},
	}, nil
}
