package commands

import (
	"github.com/spf13/cobra"
	"mingdos.com/template/app"
)

var serverCmd = &cobra.Command{
	Use:          "server",
	Short:        "Run reynold server",
	RunE:         serverCmdF,
	SilenceUsage: true,
}

func init() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.RunE = serverCmdF
}

func serverCmdF(command *cobra.Command, args []string) error {
	app.Hello(args[0])
	return nil
}
