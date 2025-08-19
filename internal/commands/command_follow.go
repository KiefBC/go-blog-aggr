package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/KiefBC/blog-aggr/internal/database"
	"github.com/google/uuid"
)

// HandlerFollow handles the 'follow' command.
// It takes in a URL as an argument and adds it to the user's followed feeds.
func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	url := cmd.Args[0]
	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not find feed with URL %s: %w", url, err)
	}

	createFeedFollow, err := s.Db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("could not create feed follow: %w", err)
	}

	fmt.Printf("User %s is now following feed %s\n", createFeedFollow.UserName, createFeedFollow.FeedName)
	return nil
}
