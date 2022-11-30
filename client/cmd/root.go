package cmd

import (
	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	cli := &cobra.Command{
		Use: "run-client",
		Short: "RC",
	}

	// cli.AddCommand(serverCmd())
	cli.AddCommand(clientCmd())

	return cli
}