package migrate

import (
	"fmt"
	"github.com/MigAru/poseidon/cmd/commands"
	"github.com/spf13/cobra"
)

func NewMigrate() *cobra.Command {
	var driver string
	var command string
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "run migrate migrations to sql",
		Long:  "run migrate migrations to sql | need exec with flags in package goose",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(driver, command)
		},
	}
	cmd.DisableFlagParsing = true
	return cmd
}

func init() {
	commands.Add(NewMigrate())
}
