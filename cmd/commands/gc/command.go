package gc

import (
	"context"
	"fmt"
	"github.com/MigAru/poseidon/cmd/commands"
	"github.com/MigAru/poseidon/cmd/commands/gc/providers"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	commands.Add(&cobra.Command{
		Use:   "gc",
		Short: "garbage collector",
		Long:  "run garbage collector for delete blobs, manifests, repositories files",
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
			_, cleanup, err := providers.InitializeApp(context.Background())
			if err != nil {
				panic(err)
			}

			<-ctx.Done()

			cleanup()
		},
	})
}
