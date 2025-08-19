package commands

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	user, err := s.Db.GetUser(context.Background(), s.Config.Current_user_name)
	if err != nil {
		return fmt.Errorf("could not find user %s: %w", s.Config.Current_user_name, err)
	}

	usersFollow, err := s.Db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get followed feeds for user %s: %w", s.Config.Current_user_name, err)
	}

	for _, follow := range usersFollow {
		fmt.Println(follow.FeedName)
	}

	return nil
}
