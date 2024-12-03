package cmd

import (
	"baolhq/gotem/lib"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func AddCmd() *cobra.Command {
	var profile string
	var createFlag bool
	var exclude []string

	cmd := &cobra.Command{
		Use:     "add file ...",
		Short:   "Adds file(s) to the stash for tracking",
		Example: "gotem add ~/.bashrc ~/.vimrc",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Collect all exclusion paths (split space-separated values)
			var excludePaths []string
			for _, ex := range exclude {
				// Split space-separated paths in each flag and append to excludePaths
				excludePaths = append(excludePaths, strings.Fields(ex)...)
			}

			// Print the exclude paths to verify all are being captured
			fmt.Println("Exclusion paths:", excludePaths)

			// Load configuration
			configPath := "./config.json"
			config, err := lib.LoadConfig(configPath)
			if err != nil {
				fmt.Printf("Error loading config.json: %v\n", err)
				return
			}

			// Get profile or default to "main"
			if profile == "" {
				profile = "main"
			}

			// Read global create setting, override with command flag if set
			create := config.Create
			if cmd.Flags().Changed("create") {
				create = createFlag
			}

			// Construct stash directory path
			profileConfig, exists := config.Profiles[profile]
			if !exists {
				profileConfig = lib.Profile{
					Dotpath: new(string),
				}
			}

			stashDir := config.Dotpath
			if profileConfig.Dotpath != nil {
				stashDir = *profileConfig.Dotpath
			}
			stashDir = filepath.Join(stashDir, profile)

			if create {
				if err := os.MkdirAll(stashDir, os.ModePerm); err != nil {
					fmt.Printf("Error creating profile directory %s: %v\n", stashDir, err)
					return
				}
			}

			// Convert exclude paths to a map for faster lookups
			excludeMap := make(map[string]struct{})
			for _, path := range excludePaths {
				absPath, err := lib.ExpandPath(path)
				if err != nil {
					fmt.Printf("Error expanding exclude path %s: %v\n", path, err)
					continue
				}

				// Debug: Verify that both exclusion paths are captured
				fmt.Printf("Expanded exclusion path: %s\n", absPath)

				excludeMap[absPath] = struct{}{}
			}

			// Debug: Show all exclusions
			fmt.Println("All exclusions:", excludeMap)

			for _, file := range args {
				// Expand leading '~' or '.' to absolute path
				srcPath, err := lib.ExpandPath(file)
				if err != nil {
					fmt.Printf("Error expanding path %s: %v\n", file, err)
					continue
				}

				// Skip excluded paths
				if _, excluded := excludeMap[srcPath]; excluded {
					fmt.Printf("Skipping excluded path: %s\n", srcPath)
					continue
				}

				info, err := os.Stat(srcPath)
				if err != nil {
					fmt.Printf("File %s does not exist. Skipping...\n", srcPath)
					continue
				}

				// Adjust the destination path based on keepdot setting
				baseName := filepath.Base(srcPath)
				if !config.Keepdot && strings.HasPrefix(baseName, ".") {
					baseName = strings.TrimPrefix(baseName, ".")
				}
				dstPath := filepath.Join(stashDir, baseName)

				// Copy file or directory
				if info.IsDir() {
					err = lib.CopyDir(srcPath, dstPath, excludeMap)
					if err != nil {
						fmt.Printf("Error copying directory %s: %v\n", srcPath, err)
						continue
					}
				} else {
					err = lib.CopyFile(srcPath, dstPath)
					if err != nil {
						fmt.Printf("Error copying file %s: %v\n", srcPath, err)
						continue
					}
				}

				// Update the configuration
				err = lib.UpdateConfig(config, profile, srcPath, baseName, info.IsDir())
				if err != nil {
					fmt.Printf("Error updating config for %s: %v\n", srcPath, err)
				}

				fmt.Printf("Successfully added %s to stash %s.\n", srcPath, profile)
			}

			// Save the updated configuration
			err = lib.SaveConfig(configPath, config)
			if err != nil {
				fmt.Printf("Error saving updated config: %v\n", err)
			}
		},
	}

	// Add --profile, --create and --exclude flag
	cmd.Flags().StringVarP(&profile, "profile", "p", "main", "Specify the profile to use")
	cmd.Flags().BoolVarP(&createFlag, "create", "c", true, "Whether the program should create new files or directories")
	cmd.Flags().StringSliceVarP(&exclude, "exclude", "e", nil, "Specify space-separated paths to exclude from addition")

	return cmd
}
