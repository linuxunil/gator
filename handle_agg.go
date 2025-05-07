package main

import (
	"context"
	"fmt"

	"github.com/linuxunil/gator/internal/database"
)

func handleAgg(st *state, cmd command) error {
	// FIXME: handle any url
	// NOTE: For now we just use static url for testing
	url := "https://www.wagslane.dev/index.xml"
	//checkArgs(len(cmd.args), 1)
	// url := cmd.args[0]
	rss, err := FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(rss)
	return nil
}

func handleAddFeed(st *state, cmd command) error {
	checkArgs(len(cmd.args), 2)
	usr, err := st.db.GetUser(context.Background(), st.cfg.Username)
	checkErr(err)
	f, err := st.db.CreateFeed(context.Background(),
		database.CreateFeedParams{Name: cmd.args[0], Url: cmd.args[1], UserID: usr.ID})
	checkErr(err)
	fmt.Println(f)
	return nil
}
