package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tracked file(s).",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing..")
		},
	}
}
