package migrate

import (
	"database/sql"
	"fmt"
	"github.com/MigAru/poseidon/cmd/commands"
	"github.com/MigAru/poseidon/cmd/commands/migrate/migrations"
	_ "github.com/mattn/go-sqlite3"
	goose "github.com/pressly/goose/v3"
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
			conn, err := sql.Open(cfg.Driver, cfg.DSN)
			if err != nil {
				panic(err)
			}
			fmt.Println(args[1:])
			goose.SetBaseFS(migrations.Migrations)
			if err := goose.SetDialect(cfg.Driver); err != nil {
				panic(err)
			}
			if err := goose.Run(
				args[0],
				conn,
				"schemas/"+cfg.Driver,
				args[1:]...,
			); err != nil {
				panic(err)
			}

		},
	}
	cmd.DisableFlagParsing = true
	return cmd
}

func init() {
	commands.Add(newCommand())
}
