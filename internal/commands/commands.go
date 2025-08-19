package commands

import "fmt"

// A collection of commands mapped by their names to their handler functions
type Commands struct {
	Commands map[string]func(*State, Command) error
}

// A command definition with its name, handler, usage, and description
type CommandDefinition struct {
	Name        string
	Handler     func(*State, Command) error
	Usage       string
	Description string
}

// A command with a name and arguments
type Command struct {
	Name string
	Args []string
}

// Register a command with its name and handler function
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Commands[name] = f
}

// Run a command by looking up its handler and executing it
func (c *Commands) Run(s *State, cmd Command) error {
	if handler, exists := c.Commands[cmd.Name]; exists {
		return handler(s, cmd)
	}
	return fmt.Errorf("command '%s' not found", cmd.Name)
}

// GetCommands returns a map of command definitions
func GetCommands() map[string]CommandDefinition {
	return map[string]CommandDefinition{
		CMD_LOGIN: {
			Name:        CMD_LOGIN,
			Handler:     HandlerLogin,
			Usage:       "login <username>",
			Description: "Log in as a user. If the user does not exist, it will return an error.",
		},
		CMD_REGISTER: {
			Name:        CMD_REGISTER,
			Handler:     HandlerRegister,
			Usage:       "register <username>",
			Description: "Register a new user. If the user already exists, it will return an error.",
		},
		CMD_RESET: {
			Name:        CMD_RESET,
			Handler:     HandlerReset,
			Usage:       "reset",
			Description: "Reset the current user. This will remove the current user from the config.",
		},
		CMD_USERS: {
			Name:        CMD_USERS,
			Handler:     HandlerUsers,
			Usage:       "users",
			Description: "List all users. The current user will be marked with an asterisk.",
		},
		CMD_HELP: {
			Name:        CMD_HELP,
			Handler:     HandlerHelp,
			Usage:       "help [command] or help",
			Description: "Show help for a specific command or list all commands if no command is specified.",
		},
		CMD_AGG: {
			Name:        CMD_AGG,
			Handler:     HandlerAgg,
			Usage:       "agg <url>",
			Description: "Fetch and display RSS feed content from the provided URL.",
		},
		CMD_ADDFEED: {
			Name:        CMD_ADDFEED,
			Handler:     HandlerAddFeed,
			Usage:       "addfeed <name> <url>",
			Description: "Add a new RSS feed with the given name and URL.",
		},
		CMD_FEEDS: {
			Name:        CMD_FEEDS,
			Handler:     HandlerFeeds,
			Usage:       "feeds",
			Description: "List all RSS feeds. The current feed will be marked with an asterisk.",
		},
	}
}

// GetUsage returns the usage string for a given command name
func (c *Commands) GetUsage(commandName string) string {
	for _, cmdDef := range GetCommands() {
		if cmdDef.Name == commandName {
			return cmdDef.Usage
		}
	}
	return fmt.Sprintf(BAD_CMD, commandName)
}

// NewCommands initializes a new Commands instance and registers all command handlers
func NewCommands() *Commands {
	commands := &Commands{
		Commands: make(map[string]func(*State, Command) error),
	}

	commandDefs := GetCommands()

	for _, cmdDef := range commandDefs {
		commands.Register(cmdDef.Name, cmdDef.Handler)
	}

	return commands
}
