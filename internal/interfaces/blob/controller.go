package blob

import (
	"github.com/MigAru/poseidon/pkg/http"
)

type Controller interface {
	Upload(ctx http.Context) error
	GetUpload(ctx http.Context) error
	CreateUpload(ctx http.Context) error
	DeleteUpload(ctx http.Context) error
	Get(ctx http.Context) error
}
