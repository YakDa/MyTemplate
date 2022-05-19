package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"mingdos.com/template/model"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	RunE:  versionCmdF,
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}

func versionCmdF(command *cobra.Command, args []string) error {
	fmt.Printf("Version: %s\n", model.CurrentVersion)
	return nil
}
