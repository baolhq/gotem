package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func RestoreCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem restore ~/gotem.tar.gz ~/gotem",
		Use:     "restore source [destination]",
		Short:   "Restore a got'em backup.",
		Args:    cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 {
				src := args[0]
				fmt.Printf("Restoring archive %s...\n", src)
			} else {
				src := args[0]
				dst := args[1]
				fmt.Printf("Restoring archive %s into %s...\n", src, dst)
			}
		},
	}
}
