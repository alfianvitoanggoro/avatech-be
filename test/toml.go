package test

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Struct matching the TOML structure
type config struct {
	App struct {
		Version string `toml:"version"`
	} `toml:"app"`
}

func Toml() {
	// Read file .config.toml
	data, err := os.ReadFile(".config.toml")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	fmt.Printf("Data: %s\n", data)

	// Parse TOML
	var config config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing TOML:", err)
		return
	}

	// Cetak versi
	fmt.Println("Version:", config.App.Version)
}
