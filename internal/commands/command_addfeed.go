package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/KiefBC/blog-aggr/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	currentUser, err := s.Db.GetUser(context.Background(), s.Config.Current_user_name)
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	createFeed, err := s.Db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      feedName,
			Url:       feedURL,
			UserID:    currentUser.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}

	fmt.Printf("Feed %s created with ID %s\n", feedName, createFeed.ID)

	return nil
}
