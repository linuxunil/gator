package main

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
	if err := c.registeredCommands[cmd.name](st, cmd); err != nil {
		return err
	}
	return nil
}
