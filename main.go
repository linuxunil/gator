package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/linuxunil/gator/internal/config"
	"github.com/linuxunil/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Note enough arguments")
		os.Exit(1)
	}
	var cfg state
	config := config.Read()
	cfg.cfg = &config
	db, err := sql.Open("postgres", cfg.cfg.DbUrl)
	cfg.db = database.New(db)
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	login := command{name: "login", args: os.Args[1:]}
	register := command{name: "register", args: os.Args[1:]}
	cmds.retisger("login", handleLogin)
	cmds.retisger("register", handleRegister)
	cmds
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
