package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func UpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [file ...]",
		Short: "Update the stash with local file(s).",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Updating all files...")
			} else {
				for _, file := range args {
					fmt.Printf("Updating %s...\n", file)
				}
			}
		},
	}
}
