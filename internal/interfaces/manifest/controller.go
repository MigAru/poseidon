package manifest

import (
	"poseidon/pkg/http"
)

type Controller interface {
	Get(ctx http.Context) error
	Create(ctx http.Context) error
	Delete(ctx http.Context) error
}
