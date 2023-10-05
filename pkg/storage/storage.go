package storage

type ST interface {
	Get(key string) (string, error)
	Update(key, value string) error
	Create(key, value string) error
	Delete(key string) error
}
