package providers

import (
	"context"
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/google/wire"
)

var servicesSet = wire.NewSet(BackendServiceProvider)

func BackendServiceProvider(ctx context.Context, server http.Server) (Backend, error) {
	server.Run(ctx)
	return Backend{}, nil
}
