package cli

// Config and Getters
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
