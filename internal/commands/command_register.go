package commands

import (
	"context"
	"fmt"
	"github.com/KiefBC/blog-aggr/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

// HandlerRegister creates a new user in the database and sets it as the current user.
func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	username := cmd.Args[0]

	_, err := s.Db.GetUser(context.Background(), username)
	if err == nil {
		fmt.Printf("User %s already exists\n", username)
		os.Exit(1)
	}

	createUser, err := s.Db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      username,
		},
	)
	if err != nil {
		fmt.Printf("failed to create user: %v\n", err)
		os.Exit(1)
	}

	err = s.Config.SetUser(username)
	if err != nil {
		fmt.Printf("failed to set user in config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("User %s created with ID %s\n", createUser.Name, createUser.ID)

	return nil
}
