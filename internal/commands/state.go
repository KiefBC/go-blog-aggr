package commands

import (
	"github.com/KiefBC/blog-aggr/internal/config"
	"github.com/KiefBC/blog-aggr/internal/database"
)

// State holds the application state, including configuration and database queries.
type State struct {
	Config   *config.Config
	Db       *database.Queries
	Commands *Commands
}
