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

func set_config(config Config) error {
	file, err := os.OpenFile(DEFAULT_CONFIG_PATH, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Error opening file: %w", err)
	}

	if err = json.NewEncoder(file).Encode(config); err != nil {
		return fmt.Errorf("Error Encoding config: %w", err)
	}

	return nil
}

func (c *Config) GetUser() string {
	return c.Number
}

func (c *Config) SetUser(name string, number string) error {
	c.User = name
	c.Number = number

	config := Config{
		User:   c.User,
		Number: c.Number,
	}

	if err := set_config(config); err != nil {
		return fmt.Errorf("Error Setting User: %w", err)
	}

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
