package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bruhjeshhh/gatorCLI/internal/database"
	"github.com/google/uuid"
)

func addFollow(s *state, cmd command) error {

	feedurl := cmd.omfo[1]
	feed, err := s.db.GetFeedby_Url(context.Background(), feedurl)
	if err != nil {
		return fmt.Errorf("could not fetch from feed %w", err)
	}
	feedid := feed.ID

	curruser, erro := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if erro != nil {
		return fmt.Errorf("could not fetch current user %w", erro)
	}

	_, er := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    curruser.ID,
		FeedID:    feedid,
	})
	if er != nil {
		log.Fatal("couldnt follow  ", er)
	}
	fmt.Println(s.cfg.CurrentUserName)
	fmt.Println(feed.Name)
	return nil
}

func getfollwedfeedsby_User(s *state, cmd command) error {
	// fmt.Println("here1")
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldnt fetch user id")
	}
	userid := user.ID
	feedsfollowed, errir := s.db.GetFeedFollowsForUser(context.Background(), userid)
	if errir != nil {
		return fmt.Errorf("couldnt fetch feeds")
	}
	fmt.Println(feedsfollowed)
	for _, id := range feedsfollowed {
		fmt.Println("here3")
		fmt.Println("yo", id)
		feed, errrr := s.db.GetFeedby_Id(context.Background(), id)
		if errrr != nil {
			return fmt.Errorf("couldnt fetch feed name")
		}
		fmt.Println("yohooo", feed)
	}
	return nil
}

func unfollow(s *state, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldnt fetch user id")
	}
	userid := user.ID

	feed, erro := s.db.GetFeedby_Url(context.Background(), cmd.omfo[1])

	if erro != nil {
		return fmt.Errorf("couldnt fetch feed id")
	}
	feedid := feed.ID

	errrr := s.db.UnfollowFeeds(context.Background(), database.UnfollowFeedsParams{
		FeedID: feedid,
		UserID: userid,
	})
	if errrr != nil {
		return fmt.Errorf("couldm't unfollow feed")
	}
	return nil
}
