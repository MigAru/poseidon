package providers

import (
	"context"
	"github.com/MigAru/poseidon/pkg/http"
)

func ServiceProvider(_ context.Context, server http.Server) (Backend, func(), error) {
	server.Run()
	return Backend{}, func() {
		server.Shutdown()
	}, nil
}
