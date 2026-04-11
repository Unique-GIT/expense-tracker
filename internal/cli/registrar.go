package cli

type cliFunc func(s *State, c *Command) error

type commandRegistrar struct {
	// maps command name to handler
	registry map[string]cliFunc
}

func (c *commandRegistrar) Register(commandName string, handler cliFunc) {
	c.registry[commandName] = handler
}

func NewRegistrar() *commandRegistrar {
	return &commandRegistrar{
		registry: make(map[string]cliFunc),
	}
}
