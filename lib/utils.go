package lib

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
)

// IsRoot Check if the program running under root privilege
func IsRoot() bool {
	return os.Getuid() == 0
}

func ExpandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[1:]), nil
	}

	if strings.HasPrefix(path, ".") {
		return filepath.Abs(path)
	}

	return path, nil
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(srcFile)

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(dstFile)

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return dstFile.Sync()
}

func CopyDir(src, dst string) error {
	// Get properties of the source directory
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to access source directory: %w", err)
	}

	if !srcInfo.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	// Ensure the destination directory exists
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Walk through the source directory
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking the file tree: %w", err)
		}

		// Construct the destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return fmt.Errorf("error calculating relative path: %w", err)
		}
		destPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			// Create subdirectories in the destination
			if err := os.MkdirAll(destPath, info.Mode()); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", destPath, err)
			}
		} else {
			// Copy files to the destination
			if err := CopyFile(path, destPath); err != nil {
				return fmt.Errorf("failed to copy file %s to %s: %w", path, destPath, err)
			}
		}
		return nil
	})
}

// Helper function to sanitize the key
func sanitizeKey(key string) string {
	// Remove special characters like dots or slashes from the key
	return strings.ReplaceAll(key, ".", "")
}

func UpdateConfig(originalPath, stashPath, profile string) error {
	configPath := "./config.toml" // Path to your configuration file

	// Load the existing configuration
	config, err := toml.LoadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Get the base name of the file/directory and sanitize it
	baseName := filepath.Base(originalPath)
	sanitizedKey := sanitizeKey(baseName)

	// Profile section key
	profileKey := fmt.Sprintf("stash.%s", profile)

	// Add profile section if it doesn't exist
	if !config.Has(profileKey) {
		config.Set(profileKey+".dotpath", profile)
	}

	// Define the full key under the profile
	fullKey := fmt.Sprintf("%s.%s", profileKey, sanitizedKey)
	config.Set(fullKey+".local_path", originalPath)
	config.Set(fullKey+".stash_path", stashPath)

	// Save the updated configuration
	f, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(f)

	_, err = config.WriteTo(f)
	if err != nil {
		return fmt.Errorf("failed to write to config file: %w", err)
	}

	return nil
}

func Decode() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(file)

	var config Config
	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to decode TOML: %v", err)
	}

	fmt.Printf("Parsed config: %v\n", config)

	profile, exists := config.Stash.Profiles["default"]
	fmt.Println(exists)
	if exists {
		fmt.Println("Default Profile dotpath", *profile.Dotpath)
		if nvimConfig, ok := profile.Files["nvim"]; ok {

			fmt.Println("Neovim Local Path: ", nvimConfig.LocalPath)
		}
	}
}
