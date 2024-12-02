package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml"
)

func IsRoot() bool {
  return os.Getuid() == 0
}

func Decode() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

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
