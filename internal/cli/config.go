package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var DEFAULT_CONFIG_PATH = ".config.json"

// Config and Getters
type Config struct {
	User   string `json:"user"`
	Number string `json:"number"`
}

func load_config() (Config, error) {
	file, err := os.Open(DEFAULT_CONFIG_PATH)
	if err != nil {
		return Config{}, fmt.Errorf("Error opening config file: %w", err)
	}

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return Config{}, fmt.Errorf("Error decoding config file: %w", err)
	}

	return config, nil
}

func (c *Config) GetUser() string {
	return c.User
}

func (c *Config) SetUser(name string, number string) error {
	c.User = name
	c.Number = number
	return nil
}

func newConfig() *Config {
	config, err := load_config()
	if err != nil {
		log.Printf("Error getting config: %v", err)
		return &Config{}
	}

	return &Config{
		User:   config.User,
		Number: config.Number,
	}
}
