package commands

import (
	"context"
	"fmt"
	"os"
)

func HandlerUsers(s *State, cmd Command) error {
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
