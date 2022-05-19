package commands

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "reynold",
	Short: "reynold is a book keeper",
	Long:  "reynold is the name of my son, and he is a book keeper",
}

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "", "configuration file to use")
}
