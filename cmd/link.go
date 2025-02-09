package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func LinkCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "link [file ...]",
		Short:   "Create symlink in the stash to local file(s)",
    Example: "gotem link ~/.bashrc",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Linking all files...")
			} else {
				for _, file := range args {
					fmt.Printf("Linking %s...\n", file)
				}
			}
		},
	}
}
