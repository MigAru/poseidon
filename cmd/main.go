package main

import (
	"github.com/MigAru/poseidon/cmd/commands"
	_ "github.com/MigAru/poseidon/cmd/commands/api"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{}

func main() {
	commands.Register(rootCMD)
	commands.SetDefaultCommandIfNonePresent(rootCMD)
	if err := rootCMD.Execute(); err != nil {
		panic(err)
	}
}
