package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// command methods
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
func (c *commands) run(st *state, cmd command) error {
	com, ok := c.registeredCommands[cmd.name]
	if !ok {
		return fmt.Errorf("Command %v does not exist\n", cmd.name)
	}
	if err := com(st, cmd); err != nil {
		return err
	}
	return nil
}
