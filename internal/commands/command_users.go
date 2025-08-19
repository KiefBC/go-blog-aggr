package commands

import (
	"context"
	"fmt"
	"os"
)

// HandlerUsers lists all users in the database and marks the current user.
func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("failed to get users: %v\n", err)
		os.Exit(1)
	}

	for _, user := range users {
		if user.Name == s.Config.Current_user_name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
