package cli

import "github.com/Unique-GIT/expense-tracker/internal/database"

type Command struct {
	Name      string
	Arguments []string
}

func NewCommand(name string, arguments []string) *Command {
	return &Command{
		Name:      name,
		Arguments: arguments,
	}
}

// Config nad Getters
type Config struct {
	user   string
	number string
}

func (c *Config) GetUser() string {
	return c.user
}

func (c *Config) SetUser(name string, number string) error {
	c.user = name
	c.number = number
	return nil
}

func newConfig() *Config {
	return &Config{
		user:   "random_user",
		number: "238213",
	}
}

// State and Getters
type State struct {
	// db
	db     *database.Queries
	config *Config
}

func (s *State) GetDb() *database.Queries {
	return s.db
}

func (c *State) GetConfig() *Config {
	return c.config
}

func NewState(inputDb *database.Queries) *State {
	return &State{
		db:     inputDb,
		config: newConfig(),
	}
}
