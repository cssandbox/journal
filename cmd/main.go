package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cssandbox/journal"
)

func main() {
	lambda.Start(journal.Router)
}
