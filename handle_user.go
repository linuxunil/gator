package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/linuxunil/gator/internal/database"
	"os"
	"time"
)

func handleUsers(st *state, cmd command) error {
	res, err := st.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for u := range res {
		if res[u].Name == st.cfg.Username {
			fmt.Printf("* %v (current)", res[u].Name)
		}
		fmt.Printf("* %v\n", res[u].Name)
	}
	return nil
}
func handleRegister(st *state, cmd command) error {
	usr, err := st.db.CreateUser(context.Background(), database.CreateUserParams{ID: int32(uuid.New()[0]), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[1]})
	if err != nil {
		return err
	}
	st.cfg.SetUser(usr.Name)
	fmt.Printf("User %v created\n", usr.Name)
	return nil

}
func handleLogin(st *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Usage: login <username>\n")
	}
	usr, err := st.db.GetUser(context.Background(), cmd.args[1])
	if err != nil {
		return err
	}
	if usr.Name != cmd.args[1] {
		os.Exit(1)
	}
	if err := st.cfg.SetUser(cmd.args[1]); err != nil {
		return err
	}
	fmt.Println("Username set")
	return nil
}
