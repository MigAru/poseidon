package migrate

import (
	"github.com/MigAru/poseidon/cmd/commands"
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "run migrate migrations to sql",
		Long:  "run migrate migrations to sql | need exec with flags in package goose",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := newConfig()
			if cfg == nil {
				panic("in parse cfg has been error")
			}

		},
	}
	cmd.DisableFlagParsing = true
	return cmd
}

func init() {
	commands.Add(newCommand())
}
