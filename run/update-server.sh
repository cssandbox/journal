(cd ../cmd/ && env GOOS=linux GOARCH=amd64 go build -o ../dist/main)
(cd ../dist && zip -j journal.zip main)
aws lambda update-function-code --function-name journal --zip-file fileb:///Users/chai/workspace/go/journal/dist/journal.zip