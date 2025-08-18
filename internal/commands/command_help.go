package commands

import (
	"fmt"
)

// HandlerHelp displays help information for commands
func HandlerHelp(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		// Show specific command help
		specificCommand := cmd.Args[0]
		cmdDef, exists := GetCommands()[specificCommand]
		if !exists {
			return fmt.Errorf("unknown command: %s", specificCommand)
		}

		fmt.Printf("Usage:: %s\nDescription: %s\n", cmdDef.Usage, cmdDef.Description)
	} else {
		// Show all commands
		for _, cmdDef := range GetCommands() {
			fmt.Printf("Command: %s\n", cmdDef.Name)
			fmt.Printf("Usage: %s\nDescription: %s\n\n", cmdDef.Usage, cmdDef.Description)
		}
	}

	return nil
}
