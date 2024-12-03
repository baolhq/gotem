package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func StatusCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem status ~/.bashrc",
		Use:     "status [file ...]",
		Short:   "Show differences between local files and the stash.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Checking status...")
			} else {
				for _, file := range args {
					fmt.Printf("Checking status of %s...\n", file)
				}
			}
		},
	}
}
