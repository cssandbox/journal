package main

import (
	"flag"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cssandbox/journal"
)

func main() {
	envPtr := flag.String("env", "prod", "Environment (dev|prod) defaults to prod")

	if *envPtr == "dev" {
		// s := &http.Server{
		// 	Addr:           ":8080",
		// 	Handler:        myHandler,
		// 	ReadTimeout:    10 * time.Second,
		// 	WriteTimeout:   10 * time.Second,
		// 	MaxHeaderBytes: 1 << 20,
		// }
		// log.Fatal(s.ListenAndServe())
	} else {
		db := dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))
		j := journal.Journal{&journal.DynamoDBStore{DB: db}}
		lambda.Start(journal.LambdaHandler(j))
	}
}
