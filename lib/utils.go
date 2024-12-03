package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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
	defer Close(srcFile)

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer Close(dstFile)

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

// LoadConfig reads and parses the JSON configuration file.
func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &config, nil
}

// SaveConfig writes the configuration back to the JSON file.
func SaveConfig(filePath string, config *Config) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print with indentation
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

// UpdateConfig updates the configuration with new entries for a given profile.
func UpdateConfig(config *Config, profileName, srcPath, dstPath string, isDir bool) error {
	// Ensure the profile exists
	profile, exists := config.Profiles[profileName]
	if !exists {
		profile = Profile{
			Directories: make(map[string]Entry),
			Files:       make(map[string]Entry),
		}
	}

	// Sanitize key names
	key := sanitizeKey(filepath.Base(srcPath))

	// Update the appropriate map
	if isDir {
		profile.Directories[key] = Entry{
			Dst: dstPath,
			Src: srcPath,
		}
	} else {
		profile.Files[key] = Entry{
			Dst: dstPath,
			Src: srcPath,
		}
	}

	// Save the updated profile back to the configuration
	config.Profiles[profileName] = profile

	return nil
}

// Helper function to sanitize keys
func sanitizeKey(key string) string {
	return strings.ReplaceAll(key, ".", "_")
}

func Close(f *os.File) {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(f)
}

func PrettyPrint(config *Config) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Printf("Error pretty-printing config: %v\n", err)
		return
	}
	fmt.Println(string(data))
}
