package api

import (
	"context"
	"fmt"
	"github.com/MigAru/poseidon/cmd/commands"
	"github.com/MigAru/poseidon/cmd/commands/api/providers"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	commands.Add(&cobra.Command{
		Use:   commands.NameDefaultCommand,
		Short: "default command",
		Long:  "run api service | default command in docker",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background())
			go func(cancel context.CancelFunc) {
				sigs := make(chan os.Signal, 1)
				signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
				<-sigs
				fmt.Println("Terminating app...")
				cancel()
			}(cancel)
			fmt.Println("Starting service...")
			_, cleanup, err := providers.InitializeBackend(ctx)
			if err != nil {
				panic(err)
			}

			<-ctx.Done()

			cleanup()
		},
	})
}
