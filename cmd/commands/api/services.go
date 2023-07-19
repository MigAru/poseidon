package api

import (
	"github.com/MigAru/poseidon/pkg/http"
	"github.com/google/wire"
)

var servicesSet = wire.NewSet(BackendServiceProvider)

func BackendServiceProvider(server http.HttpServer) (Backend, error) {
	server.Run()
	return Backend{}, nil
}
