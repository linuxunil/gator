package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/linuxunil/gator/internal/database"
)

func handleFeeds(st *state, cmd command) error {
	feeds, err := st.db.GetFeeds(context.Background())
	checkErr(err)
	for i := range feeds {
		user, err := st.db.GetUserByID(context.Background(), feeds[i].UserID)
		checkErr(err)
		fmt.Printf("name: %v\nurl: %v\nuser:%v\n", feeds[i].Name, feeds[i].Url, user)

	}
	return nil
}
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
	usr, err := st.db.GetUserByName(context.Background(), st.cfg.Username)
	checkErr(err)
	feed, err := st.db.CreateFeed(context.Background(),
		database.CreateFeedParams{ID: uuid.New(), Name: cmd.args[0], Url: cmd.args[1], UserID: usr.ID})
	checkErr(err)
	_, err = st.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), UserID: usr.ID, FeedID: feed.ID})
	checkErr(err)
	fmt.Printf("user %v added %v", usr.Name, feed.Name)
	return nil
}
