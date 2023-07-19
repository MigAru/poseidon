package manifest

import (
	"github.com/MigAru/poseidon/pkg/http"
)

type Controller interface {
	Get(ctx http.Context) error
	Create(ctx http.Context) error
	Delete(ctx http.Context) error
}
