package main

import (
	"context"
	"fmt"
	"os"

	"github.com/linuxunil/gator/internal/database"
	"github.com/linuxunil/gator/internal/feed"
)

func handleAgg(st *state, cmd command) error {
	rss, err := feed.FetchFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println(rss)
	return nil
}

func handleAddFeed(st *state, cmd command) error {
	if len(cmd.args) < 3 {
		return fmt.Errorf("Usage: addfeed <name> <url>\n")
		os.Exit(1)
	}
	usr, err := st.db.GetUser(context.Background(), st.cfg.Username)
	if err != nil {
		return err
	}
	f, err := st.db.CreateFeed(context.Background(),
		database.CreateFeedParams{Name: cmd.args[0], Url: cmd.args[1], UserID: usr.ID})
	if err != nil {
		return err
	}
	fmt.Println(f)
	return nil
}
