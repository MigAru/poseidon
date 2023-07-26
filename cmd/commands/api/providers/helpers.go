package providers

import (
	"github.com/MigAru/poseidon/pkg/registry/hasher"
	"github.com/google/wire"
)

var helpersSet = wire.NewSet(
	hasher.NewHasher,
)
