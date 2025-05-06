package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/linuxunil/gator/internal/database"
)

func handleRegister(st *state, cmd command) error {
	usr, err := st.db.CreateUser(context.Background(), database.CreateUserParams{ID: int32(uuid.New()[0]), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[2]})
	st.cfg.SetUser(usr.Name)
	if err != nil {
		return err
		os.Exit(1)
	}
	return nil

}
func handleLogin(st *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Usage: login <username>\n")
	}
	if err := st.cfg.SetUser(cmd.args[1]); err != nil {
		return err
	}
	fmt.Println("Username set")
	return nil
}
