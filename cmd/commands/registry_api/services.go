package registry_api

import (
	"github.com/google/wire"
	"poseidon/pkg/http"
)

var servicesSet = wire.NewSet(BackendServiceProvider)

func BackendServiceProvider(server http.HttpServer) (Backend, error) {
	server.Run()
	return Backend{}, nil
}
