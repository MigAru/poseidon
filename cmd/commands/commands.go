package commands

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	Commands           []*cobra.Command
	NameDefaultCommand = "app"
)

func Add(command *cobra.Command) {
	Commands = append(Commands, command)
}

func Register(rootCommand *cobra.Command) {
	for _, command := range Commands {
		rootCommand.AddCommand(command)
	}
}

func SetDefaultCommandIfNonePresent(rootCmd *cobra.Command) {
	cmd, _, err := rootCmd.Find(os.Args[1:])
	if err != nil || cmd.Use == "" {
		args := append([]string{NameDefaultCommand}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

}
