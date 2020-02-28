package journal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DynamoDBStore DynamoDB implementation of Store
type DynamoDBStore struct {
	DB *dynamodb.DynamoDB
}

//GetItems implementation for Store interface
func (store *DynamoDBStore) GetItems() (entries []*Entry, err error) {
	// Prepare the input for the query.
	input := &dynamodb.ScanInput{
		TableName: aws.String("Journal"),
	}

	// Retrieve the item from DynamoDB. If no matching item is found
	// return nil.
	result, err := store.DB.Scan(input)
	if err != nil {
		return nil, err
	}
	if result.Items == nil {
		return nil, nil
	}

	for _, i := range result.Items {
		en := new(Entry)
		err = dynamodbattribute.UnmarshalMap(i, en)
		if err != nil {
			return nil, err
		}
		entries = append(entries, en)
	}

	return entries, nil
}

//GetItem implementation of Store interface
func (store *DynamoDBStore) GetItem(UUID string) (*Entry, error) {
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
	result, err := store.DB.GetItem(input)
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
	en := new(Entry)
	err = dynamodbattribute.UnmarshalMap(result.Item, en)
	if err != nil {
		return nil, err
	}

	return en, nil
}

//PutItem implementation for Store interface
func (store *DynamoDBStore) PutItem(en *Entry) error {
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

	_, err := store.DB.PutItem(input)
	return err
}
