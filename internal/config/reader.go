package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

var (
	cfg *Config
	// ErrUnknownFileExtension is returned by the Parse function
	// when the file extension is not allowed for configuration
	ErrUnknownFileExtension = errors.New("unknown file extension")
)

func Parse(path string, cfg *Config) error {
	switch fileExtension(path) {
	case "yaml":
		return parseYAML(path, cfg)
	default:
		return ErrUnknownFileExtension
	}
}

func ReadEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}

func SetConfig(c *Config) {
	cfg = c
}

func fileExtension(path string) string {
	s := strings.Split(path, ".")
	return s[len(s)-1]
}

// parseYAML parses yaml config file into Config
func parseYAML(path string, cfg *Config) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		err = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}
