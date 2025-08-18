package commands

import (
	"context"
	"fmt"
	"os"
)

// HandlerReset resets the database by dropping all tables and recreating them.
func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("Usage: %s", s.Commands.GetUsage(cmd.Name))
	}

	err := s.Db.Reset(context.Background())
	if err != nil {
		fmt.Printf("failed to reset database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Database reset successfully")
	return nil
}
