package cmd

import (
	"baolhq/gotem/lib"
	"fmt"

	"github.com/spf13/cobra"
)

func TestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "gotem test",
		Short: "Test program",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := lib.LoadConfig("./config.json")
			if err != nil {
				fmt.Printf("Failed to load config.json: %v", err)
			}
			lib.PrettyPrint(config)
		},
	}
}
