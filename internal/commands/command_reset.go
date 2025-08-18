package commands

import (
	"context"
	"fmt"
	"os"
)

func HandleReset(s *State, cmd Command) error {
	err := s.Db.Reset(context.Background())
	if err != nil {
		fmt.Printf("failed to reset database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database reset successfully")
	return nil
}
