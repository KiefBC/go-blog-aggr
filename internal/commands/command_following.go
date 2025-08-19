package commands

import (
	"context"
	"fmt"

	"github.com/KiefBC/blog-aggr/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	usersFollow, err := s.Db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get followed feeds for user %s: %w", user.Name, err)
	}

	for _, follow := range usersFollow {
		fmt.Println(follow.FeedName)
	}

	return nil
}
