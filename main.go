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

func checkArgs(actual int, expected int) {
	if actual < expected {
		fmt.Printf("Expected %v arguments given %v\n", expected, actual)
		os.Exit(1)
	}
}
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func main() {
	// Check that enough arguments were given
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	// Set program state
	var cfg state
	config := config.Read()
	cfg.cfg = &config
	// Database setup
	// Connect to database
	db, err := sql.Open("postgres", cfg.cfg.DbUrl)
	checkErr(err)
	// Set up go bindings for database query (sqlc)
	cfg.db = database.New(db)

	// Someplace to keep all our commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	// Register our commands for use
	cmds.register("login", handleLogin)
	cmds.register("register", handleRegister)
	cmds.register("reset", handleReset)
	cmds.register("users", handleUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", handleAddFeed)
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", handleFollow)
	cmds.register("following", handleFollowing)
	// Grab the user provided command and it's arguments
	cmd := os.Args[1]
	args := os.Args[2:]
	// Run the command
	err = cmds.run(&cfg, command{name: cmd, args: args})
	checkErr(err)

}
