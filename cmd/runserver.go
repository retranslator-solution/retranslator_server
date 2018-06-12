package cmd

import (
	"github.com/retranslator-solution/retranslator_server/application"
	"github.com/retranslator-solution/retranslator_server/configs"
	"github.com/retranslator-solution/retranslator_server/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runserverCmd represents the runserver command
var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run Retranslator server",
	Run: func(cmd *cobra.Command, args []string) {
		config := configs.NewConfig(viper.GetViper())
		app := application.NewApplication(config)

		defer app.Storage.Close()

		server.RunServer(app)
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)
}
