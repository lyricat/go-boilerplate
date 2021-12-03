package migrate

import (
	"github.com/fox-one/pkg/store/db"
	"github.com/lyricat/go-boilerplate/config"
	"github.com/spf13/cobra"
)

func NewCmdMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"setdb"},
		Short:   "migrate database tables",
		Run: func(cmd *cobra.Command, args []string) {
			database := db.MustOpen(config.C().DB)
			cmd.Println(config.C().DB)
			defer database.Close()

			if err := db.Migrate(database); err != nil {
				cmd.PrintErrln("migrate tables", err)
				return
			}
			cmd.Println("migrate done")
		},
	}

	return cmd
}
