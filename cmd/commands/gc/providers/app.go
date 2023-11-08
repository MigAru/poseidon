package providers

import (
	"context"
	"github.com/MigAru/poseidon/internal/gc"
)

func AppProvider(_ context.Context, gc *gc.GC) (App, func(), error) {
	go gc.Start()
	return App{}, func() {
	}, nil
}
