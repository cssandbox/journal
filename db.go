package journal

// Store For Journal entries
type Store interface {
	GetItems() ([]*Entry, error)
	GetItem(string) (*Entry, error)
	PutItem(*Entry) error
}
