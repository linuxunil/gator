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
	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", handleAddFeed)
	switch os.Args[1] {
	case "addfeed":
		if err := cmds.run(&cfg, command{name: "addfeed", args: os.Args[1:]}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "agg":
		if err := cmds.run(&cfg, command{name: "agg", args: os.Args[1:]}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "users":
		if err := cmds.run(&cfg, command{name: "users", args: os.Args[1:]}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "reset":
		if err := cmds.run(&cfg, command{name: "reset", args: os.Args[1:]}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "login":
		if err := cmds.run(&cfg, command{name: "login", args: os.Args[1:]}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "register":
		if err := cmds.run(&cfg, command{name: "register", args: os.Args[1:]}); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
