package journal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func getItems() (entries []*entry, err error) {
	// Prepare the input for the query.
	input := &dynamodb.ScanInput{
		TableName: aws.String("Journal"),
	}

	// Retrieve the item from DynamoDB. If no matching item is found
	// return nil.
	result, err := db.Scan(input)
	// result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Items == nil {
		return nil, nil
	}

	for _, i := range result.Items {
		en := new(entry)
		err = dynamodbattribute.UnmarshalMap(i, en)
		entries = append(entries, en)
	}
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func getItem(UUID string) (*entry, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Journal"),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {
				S: aws.String(UUID),
			},
		},
	}

	// Retrieve the item from DynamoDB. If no matching item is found
	// return nil.
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type
	// map[string]*AttributeValue. We can use the UnmarshalMap helper
	// to parse this straight into the fields of a struct. Note:
	// UnmarshalListOfMaps also exists if you are working with multiple
	// items.
	en := new(entry)
	err = dynamodbattribute.UnmarshalMap(result.Item, en)
	if err != nil {
		return nil, err
	}

	return en, nil
}

func putItem(en *entry) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Journal"),
		Item: map[string]*dynamodb.AttributeValue{
			"UUID": {
				S: aws.String(en.UUID),
			},
			"Title": {
				S: aws.String(en.Title),
			},
			"Body": {
				S: aws.String(en.Body),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
