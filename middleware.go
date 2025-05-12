package main

import (
	"context"

	"github.com/linuxunil/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUserByName(context.Background(), s.cfg.Username)
		checkErr(err)
		handler(s, c, user)
		return nil
	}
}
