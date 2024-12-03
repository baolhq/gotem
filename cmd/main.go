package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:     "gotem",
		Short:   "Go-based Tool for Efficient Management",
		Long:    "Got'em is a simple program to manage your dotfiles",
		Version: "0.0.1",
	}

	rootCmd.AddCommand(
		AddCmd(),
		RemoveCmd(),
		LinkCmd(),
		UnlinkCmd(),
		StatusCmd(),
		ListCmd(),
		UpdateCmd(),
		InstallCmd(),
		BackupCmd(),
		RestoreCmd(),
		TestCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
