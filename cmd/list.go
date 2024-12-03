package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tracked file(s)",
    Example: "gotem list --profile home",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing..")
		},
	}
}
