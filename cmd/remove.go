package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RemoveCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "remove [file ...]",
		Short:   "Removes tracked file(s) from the stash",
    Example: "gotem remove ~/.bashrc",
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
