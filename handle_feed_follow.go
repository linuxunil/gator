package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/linuxunil/gator/internal/database"
)

func handleFollow(st *state, cmd command, usr database.User) error {
	feed, err := st.db.GetFeedByURL(context.Background(), cmd.args[0])
	checkErr(err)
	_, err = st.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), UserID: usr.ID, FeedID: feed.ID})
	checkErr(err)
	fmt.Printf("feed: %v, followed by %v\n", feed.Name, usr.Name)
	return nil

}
func handleFollowing(st *state, cmd command, usr database.User) error {
	feeds, err := st.db.GetFeedFollowsForUser(context.Background(), usr.ID)
	checkErr(err)
	fmt.Printf("User %v  logged in user %v,\n", usr.Name, st.cfg.Username)
	for i := range feeds {
		fmt.Printf("feed: %v\n", feeds[i].FeedName)
	}
	return nil
}

func handleUnfollow(st *state, cmd command, usr database.User) error {
	feed, err := st.db.GetFeedByURL(context.Background(), cmd.args[0])
	err = st.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: usr.ID, FeedID: feed.ID})
	checkErr(err)
	return nil
}
