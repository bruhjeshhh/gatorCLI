package main

import (
	"context"
	_ "database/sql"
	"fmt"
	"time"

	"github.com/bruhjeshhh/gatorCLI/internal/database"
	"github.com/google/uuid"
)

func addFeed(s *state, cmd command) error {
	fmt.Print(cmd.omfo)
	if len(cmd.omfo) != 3 {
		return fmt.Errorf("not enough arguements")
	}
	// _, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return fmt.Errorf("could not fetch %w", err)
	// }

	username := s.cfg.CurrentUserName
	user, _ := s.db.GetUser(context.Background(), username)

	_, erro := s.db.Feed_intoDb(context.Background(), database.Feed_intoDbParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.omfo[1],
		Url:       cmd.omfo[2],
		UserID:    user.ID,
	})

	if erro != nil {
		return fmt.Errorf("db couldnt update %w", erro)
	}
	fmt.Println("done")

	feedurl := cmd.omfo[2]
	feed, err := s.db.GetFeedby_Url(context.Background(), feedurl)
	if err != nil {
		return fmt.Errorf("could not fetch from feed %w", err)
	}
	feedid := feed.ID
	curruser, erro := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	_, e := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    curruser.ID,
		FeedID:    feedid,
	})
	if e != nil {
		return fmt.Errorf("error following feed  %w", e)
	}
	return nil
}

func fetchFeeds(s *state, cmd command) error {
	feed, errrr := s.db.GetFeed(context.Background())
	if errrr != nil {
		return fmt.Errorf("couldnt fetch")
	}
	fmt.Println(feed)
	return nil
}
