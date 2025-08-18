package commands

import (
	"context"
	"fmt"
	"os"
)

// HandlerLogin sets the current user based on the provided username argument.
func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("login command requires a username argument")
	}

	user, err := s.Db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Println("user not found:", err)
		os.Exit(1)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("Setting current user to %s\n", s.Config.Current_user_name)
	return nil
}
