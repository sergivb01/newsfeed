package newsfeed

type Storage interface {
	Put([]Item) error
	Get() ([]Item, error)
}
