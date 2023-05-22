package digest

type Repository interface {
	Get(project, name string) ([]byte, error)
	Create(project, name string, data []byte) error
	Exist(project, name string) error
	Delete(project, name string) error
}
