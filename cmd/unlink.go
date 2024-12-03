package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func UnlinkCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem unlink",
		Use:     "unlink [file ...]",
		Short:   "Remove symlink from the stash that point to local file(s)",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Unlinking all files?")
			} else {
				for _, file := range args {
					fmt.Printf("Unlinking %s...\n", file)
				}
			}
		},
	}
}
