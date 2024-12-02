package main

import (
	"baolhq/gotem/lib"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gotem",
		Short: "Go-based Tool for Efficient Management",
		Long:  "Got'tem is a simple program to manage your dotfiles",
	}

	commands := []*cobra.Command{
    lib.ImportCmd(),
    lib.TestCmd(),
	}

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
