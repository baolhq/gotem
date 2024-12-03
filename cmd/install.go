package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func InstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gotem install [file ...]",
		Short: "Install files from the stash into your machine",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Installing all files...")
			} else {
				for _, file := range args {
					fmt.Printf("Installing %s...\n", file)
				}
			}
		},
	}
}
