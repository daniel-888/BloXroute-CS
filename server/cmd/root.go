package cmd

import (
	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	cli := &cobra.Command{
		Use: "run-server",
		Short: "CS",
	}

	cli.AddCommand(serverCmd())
	// cli.AddCommand(clientCmd())

	return cli
}