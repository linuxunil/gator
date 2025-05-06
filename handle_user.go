package main

import (
	"fmt"
)

func handleLogin(st *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Usage: login <username>\n")
	}
	if err := st.config.SetUser(cmd.args[1]); err != nil {
		return err
	}
	fmt.Println("Username set")
	return nil
}
