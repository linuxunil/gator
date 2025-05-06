package main

import (
	"fmt"
	"os"

	"github.com/linuxunil/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Note enough arguments")
		os.Exit(1)
	}

	var curConfig state
	config := config.Read()
	curConfig.config = &config
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	login := command{name: "login", args: os.Args[1:]}
	cmds.retisger("login", handleLogin)
	err := cmds.run(&curConfig, login)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
