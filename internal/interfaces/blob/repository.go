package blob

//для не структурированного хранения слоев и для дальнейшего создания digest(более структурированного хранения слоев)
type Repository interface {
	Get(name string) ([]byte, error)
	Create(name string, data []byte) error
	Exist(name string) bool
	Delete(name string) error
}
