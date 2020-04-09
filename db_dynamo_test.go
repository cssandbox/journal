package journal

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Test_marshalledEntry(t *testing.T) {
	type args struct {
		en *Entry
	}
	tests := []struct {
		name string
		args args
		want map[string]*dynamodb.AttributeValue
	}{
		{"First", args{&Entry{
			Title: "Covid-19 April/7",
			Text:  "Worked on this Journal App. Draw each day with the Miles. Draw each day with Pallu and Akki.",
			Sections: []Section{
				{
					Title: "Journal App",
					Text:  "Updated the domain model",
				},
				{
					Title: "Draw with Miles",
					Text:  "None had time today, so it got skipped",
				},
				{
					Title: "Draw with Akki",
					Text:  "Today's theme was \"Fluid in Motion\"",
				},
			},
		}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := marshalledEntry(tt.args.en); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("marshalledEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
