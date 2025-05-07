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

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
	check(err)
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
	cmd := os.Args[1]
	args := os.Args[2:]
	err = cmds.run(&cfg, command{name: cmd, args: args})
	check(err)

}
