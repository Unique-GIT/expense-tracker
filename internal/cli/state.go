package cli

import "github.com/Unique-GIT/expense-tracker/internal/database"

type Command struct {
	Name      string
	Arguments []string
}

// Config nad Getters
type Config struct {
	user string
}

func (c *Config) GetUser() string {
	return c.user
}

func newConfig() *Config {
	return &Config{
		user: "random_user",
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
