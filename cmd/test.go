package cmd

import (
	"baolhq/gotem/lib"
	"github.com/spf13/cobra"
)

func TestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test program",
		Run: func(cmd *cobra.Command, args []string) {
			lib.Decode()
		},
	}
}
