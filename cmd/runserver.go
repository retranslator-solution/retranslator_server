package cmd

import (
	"github.com/spf13/cobra"
	"github.com/retranslator-solution/retranslator_server/server"
)

// runserverCmd represents the runserver command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run Retranslator server",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)

}
