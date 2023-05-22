package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"poseidon/cmd/commands"
	"poseidon/cmd/registry_api"
	"poseidon/internal/consts"
	"syscall"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "test task service",
		Usage: "Run test task service",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    consts.DebugMode,
				EnvVars: []string{consts.DebugMode},
				Value:   true,
			},
			&cli.IntFlag{
				Name:    consts.ServerTimeout,
				EnvVars: []string{consts.ServerTimeout},
				Value:   15,
			},
			&cli.StringFlag{
				Name:    consts.ServerPort,
				EnvVars: []string{consts.ServerPort},
				Value:   ":8000",
			},
		},
		Action: func(c *cli.Context) error {
			ctx, cancel := context.WithCancel(context.Background())
			go func() {
				sigs := make(chan os.Signal, 1)
				signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
				<-sigs
				fmt.Println("Terminating app...")
				cancel()
			}()
			_, cleanup, err := registry_api.InitializeBackend(c)

			<-ctx.Done()

			if err != nil {
				panic(err)
			}
			cleanup()
			return nil
		},
	}
	app.Commands = commands.Commands
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
