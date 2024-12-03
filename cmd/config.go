package cmd

import (
	"baolhq/gotem/lib"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

func ConfigCmd() *cobra.Command {
	var profile string
	var setOptions map[string]string
	var unsetOptions []string

	cmd := &cobra.Command{
		Use:     "gotem config",
		Short:   "Manage got'em configurations",
		Example: "gotem config --set backup=true",
		Run: func(cmd *cobra.Command, args []string) {
			// Load configuration
			configPath := "./config.json"
			config, err := lib.LoadConfig(configPath)
			if err != nil {
				fmt.Printf("Error loading config.json: %v\n", err)
				return
			}

			// Handle setting options
			if len(setOptions) > 0 {
				profileName := "main"
				if profile != "" {
					profileName = profile
				}

				// Ensure the profile exists
				if _, exists := config.Profiles[profileName]; !exists {
					config.Profiles[profileName] = lib.Profile{
						Directories: make(map[string]lib.Entry),
						Files:       make(map[string]lib.Entry),
					}
				}

				profileConfig := config.Profiles[profileName]

				for key, value := range setOptions {
					var parsedVal interface{}
					if value == "true" || value == "false" {
						parsedVal = value == "true"
					} else if intVal, err := json.Number(value).Int64(); err == nil {
						parsedVal = intVal
					} else {
						parsedVal = value
					}

					switch key {
					case "backup":
						boolVal := parsedVal.(bool)
						profileConfig.Backup = &boolVal
					case "create":
						boolVal := parsedVal.(bool)
						profileConfig.Create = &boolVal
					case "dotpath":
						strVal := parsedVal.(string)
						profileConfig.Dotpath = &strVal
					case "uselink":
						boolVal := parsedVal.(bool)
						profileConfig.Uselink = &boolVal
					case "keepdot":
						boolVal := parsedVal.(bool)
						profileConfig.Keepdot = &boolVal
					default:
						fmt.Printf("Unknown key: %s\n", key)
					}
				}

				config.Profiles[profileName] = profileConfig
			}

			// Handle unsetting options
			if len(unsetOptions) > 0 {
				if profile == "" {
					fmt.Println("Unset operation requires a specific profile. Use --profile.")
					return
				}

				profileConfig, exists := config.Profiles[profile]
				if !exists {
					fmt.Printf("Profile %s not found.\n", profile)
					return
				}

				for _, key := range unsetOptions {
					switch key {
					case "backup":
						profileConfig.Backup = nil
					case "create":
						profileConfig.Create = nil
					case "dotpath":
						profileConfig.Dotpath = nil
					case "uselink":
						profileConfig.Uselink = nil
					case "keepdot":
						profileConfig.Keepdot = nil
					default:
						fmt.Printf("Unknown key: %s\n", key)
					}
				}

				config.Profiles[profile] = profileConfig
			}

			// Save updated configuration
			if err := lib.SaveConfig(configPath, config); err != nil {
				fmt.Printf("Error saving config.json: %v\n", err)
				return
			}

			// Display configurations
			if profile != "" {
				if profileConfig, exists := config.Profiles[profile]; exists {
					// Display only the specific profile
					fmt.Printf("Configurations for \"%s\":\n", profile)
					data, err := json.MarshalIndent(profileConfig, "", "  ")
					if err != nil {
						fmt.Printf("Error formatting profile: %v\n", err)
						return
					}
					fmt.Println(string(data))
				} else {
					fmt.Printf("Profile \"%s\" not found.\n", profile)
				}
			} else {
				// Display full configuration
				fmt.Println("Global configurations:")
				lib.PrettyPrint(config)
			}
		},
	}

	// Add flags for --profile, --set, and --unset
	cmd.Flags().StringVarP(&profile, "profile", "p", "", "Specify the profile to view or modify")
	cmd.Flags().StringToStringVarP(&setOptions, "set", "s", nil, "Set configuration options (e.g., --set key=value)")
	cmd.Flags().StringArrayVarP(&unsetOptions, "unset", "u", nil, "Unset a configuration option")

	return cmd
}

