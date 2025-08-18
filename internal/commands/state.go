package commands

import (
	"github.com/KiefBC/blog-aggr/internal/config"
	"github.com/KiefBC/blog-aggr/internal/database"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}
