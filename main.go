package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/KiefBC/blog-aggr/internal/commands"
	"github.com/KiefBC/blog-aggr/internal/config"
	"github.com/KiefBC/blog-aggr/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	db, err := sql.Open("postgres", cfg.Db_url)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	state := commands.State{
		Config: cfg,
		Db:     dbQueries,
	}

	cmds := &commands.Commands{
		Commands: make(map[string]func(*commands.State, commands.Command) error),
	}
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandleReset)
	cmds.Register("users", commands.HandlerUsers)

	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: <command> [args...]\n")
		os.Exit(1)
	}

	cmd := commands.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.Run(&state, cmd)
	if err != nil {
		fmt.Printf("Error executing command '%s': %v\n", cmd.Name, err)
		os.Exit(1)
	}
}
