package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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
		fmt.Printf("name: %v\n\turl: %v\n\tuser:%v\n\tCreated at:%v\n\tUpdated at:%v\n\tFetched: %v\n", feeds[i].Name, feeds[i].Url, user, feeds[i].CreatedAt, feeds[i].UpdatedAt, feeds[i].LastFetchedAt)

	}
	return nil
}

func handleAddFeed(st *state, cmd command, usr database.User) error {
	checkArgs(len(cmd.args), 2)

	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := st.db.CreateFeed(context.Background(),
		database.CreateFeedParams{ID: uuid.New(), Name: name, Url: url, UserID: usr.ID, CreatedAt: time.Now()})
	checkErr(err)
	_, err = st.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New(), CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), UserID: usr.ID, FeedID: feed.ID})
	checkErr(err)
	fmt.Printf("user %v added %v", usr.Name, feed.Name)
	return nil
}

func handleAgg(st *state, cmd command) error {
	checkArgs(len(cmd.args), 1)
	duration, err := time.ParseDuration(cmd.args[0])
	checkErr(err)
	ticker := time.NewTicker(duration)
	fmt.Printf("Collecting data every %v\n", duration)
	for ; ; <-ticker.C {
		fmt.Println("=========Scraper tick==================")
		scrapeFeeds(st)
		fmt.Println("=========Scraper done==================")
	}
}

func scrapeFeeds(st *state) error {
	next, err := st.db.GetNextFeedToFetch(context.Background())
	st.db.MarkFeedFetched(context.Background(), next.ID)
	checkErr(err)
	feed, err := FetchFeed(context.Background(), next.Url)
	checkErr(err)
	for i := range feed.Channel.Item {
		post, err := st.db.CreatePost(context.Background(), database.CreatePostParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: sql.NullTime{Time: time.Now()}, Title: feed.Channel.Item[i].Title, Url: feed.Channel.Item[i].Link, Description: feed.Channel.Item[i].Description, FeedID: next.ID})
		checkErr(err)

		fmt.Println("===========================")
		fmt.Printf("Created %v\n", post.Title)
		fmt.Println("===========================")
	}
	return nil
}

func handleBrowse(st *state, cmd command) error {
	limit := 2
	if len(cmd.args) > 0 {
		limit, _ = strconv.Atoi(cmd.args[0])
	}
	usr, err := st.db.GetUserByName(context.Background(), st.cfg.Username)
	checkErr(err)
	posts, err := st.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: usr.ID, Limit: int32(limit)})
	for i := range posts {
		fmt.Printf("Title: %v\n\tURL: %v\n\tCreated at:%v\n\tUpdated at:%v\n", posts[i].Title, posts[i].Url, posts[i].CreatedAt, posts[i].UpdatedAt)
	}

	return nil
}
