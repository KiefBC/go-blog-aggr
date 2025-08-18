package commands

import "fmt"

type Commands struct {
	Commands map[string]func(*State, Command) error
}

type Command struct {
	Name string
	Args []string
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	if handler, exists := c.Commands[cmd.Name]; exists {
		return handler(s, cmd)
	}
	return fmt.Errorf("command '%s' not found", cmd.Name)
}
