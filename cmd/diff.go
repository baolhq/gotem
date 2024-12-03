package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func DiffCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff [file ...]",
		Short: "Show differences between local files and their stash versions",
    Example: "gotem diff ~/.bashrc",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, file := range args {
				stashFile := filepath.Join("stash", filepath.Base(file))

				if _, err := os.Stat(stashFile); os.IsNotExist(err) {
					fmt.Printf("Stash version for %s not found.\n", file)
					continue
				}

				diff := exec.Command("diff", "--color=always", stashFile, file)
				diff.Stdout = os.Stdout
				diff.Stderr = os.Stderr
				if err := diff.Run(); err != nil {
					if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() != 1 {
						fmt.Printf("Error running diff for %s: %v\n", file, err)
					}
				}
			}
		},
	}

	return cmd
}
