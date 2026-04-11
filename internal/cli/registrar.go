package cli

import "fmt"

type cliFunc func(s *State, c *Command) error

type commandRegistrar struct {
	// maps command name to handler
	registry map[string]cliFunc
}

func (c *commandRegistrar) Register(commandName string, handler cliFunc) {
	c.registry[commandName] = handler
}

func (c *commandRegistrar) Run(s *State, cmd *Command) error {
	handler, exists := c.registry[cmd.Name]
	if !exists {
		return fmt.Errorf("command doesn't exist: %s", cmd.Name)
	}

	return handler(s, cmd)
}

func NewRegistrar() *commandRegistrar {
	return &commandRegistrar{
		registry: make(map[string]cliFunc),
	}
}
