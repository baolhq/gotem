package cmd

import (
	"baolhq/gotem/lib"
	"fmt"
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

	commands := []*cobra.Command{
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
	}

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func AddCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem add ~/.bashrc ~/.vimrc",
		Use:     "add files...",
		Short:   "Adds file(s) to the stash for tracking.",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, file := range args {
				fmt.Printf("Adding file %s...\n", file)
			}
		},
	}
}

func RemoveCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem remove ~/.bashrc",
		Use:     "remove [files...]",
		Short:   "Removes tracked file(s) from the stash.",
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

func LinkCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem link",
		Use:     "link",
		Short:   "Create symlink in the stash to local file(s).",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Linking all files...")
			} else {
				for _, file := range args {
					fmt.Printf("Linking %s...\n", file)
				}
			}
		},
	}
}

func UnlinkCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem unlink",
		Use:     "unlink",
		Short:   "Remove symlink from the stash that point to local file(s).",
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

func StatusCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem status ~/.bashrc",
		Use:     "status",
		Short:   "Show differences between local files and the stash.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Checking status...")
			} else {
				for _, file := range args {
					fmt.Printf("Checking status of %s...\n", file)
				}
			}
		},
	}
}

func ListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all tracked file(s).",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Listing..")
		},
	}
}

func UpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
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

func InstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Install files from the stash into your machine.",
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

func BackupCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem backup ~/gotem.tar.gz",
		Use:     "backup [destination]",
		Short:   "Create a compressed archive of the stash.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Backing up into stash/profile...")
			} else {
				fmt.Printf("Backing up into %s...", args[0])
			}
		},
	}
}

func RestoreCmd() *cobra.Command {
	return &cobra.Command{
		Example: "gotem restore ~/gotem.tar.gz ~/gotem",
		Use:     "restore source [destination]",
		Short:   "Unarchive a got'em backup.",
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

func TestCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "test",
		Short: "Test program",
		Run: func(cmd *cobra.Command, args []string) {
			lib.Decode()
		},
	}
}
