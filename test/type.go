package test

// Struct matching the TOML structure
type App struct {
	Name    string `toml:"name" yaml:"name"`
	Version string `toml:"version" yaml:"version"`
	Debug   bool   `toml:"debug" yaml:"debug"`
}

type Config struct {
	App App `toml:"app" yaml:"app"`
}
