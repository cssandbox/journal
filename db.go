package journal

import "time"

// Store For Journal entries
type Store interface {
	GetItems() ([]*Entry, error)
	GetItem(string) (*Entry, error)
	PutItem(*Entry) error
}

// Entry for journal
type Entry struct {
	UUID        string    `json:"UUID"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Sections    []Section `json:"sections"`
	Public      bool      `json:"public"`
	CreatedDate time.Time
}

// Section of journal Entry
type Section struct {
	Title  string  `json:"title"`
	Photos []Photo `json:"photos"`
	Text   string  `json:"text"`
}

// Photo of Section
type Photo struct {
	URL string `json:"url"`
}
