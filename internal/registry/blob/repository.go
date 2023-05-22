package blob

type Repository interface {
	Get(name string) ([]byte, error)
	Create(name string, data []byte) error
	Exist(name string) bool
	Delete(name string) error
}
