package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RemoveCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem remove ~/.bashrc",
		Use:     "remove [file ...]",
		Short:   "Removes tracked file(s) from the stash.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Removing all files?")
			} else {
				for _, file := range args {
					fmt.Printf("Removing files %s...\n", file)
				}
			}
		},
	}
}
