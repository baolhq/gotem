package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func BackupCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "backup [destination]",
		Short:   "Create a compressed archive of the stash",
		Example: "gotem backup ~/gotem.tar.gz",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Backing up into stash/profile...")
			} else {
				fmt.Printf("Backing up into %s...", args[0])
			}
		},
	}
}
