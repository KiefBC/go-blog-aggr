package commands

import (
	"context"
	"fmt"
	"github.com/KiefBC/blog-aggr/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("register command requires a username argument")
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
