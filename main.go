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

	cmds := commands.NewCommands()

	state := commands.State{
		Config:   cfg,
		Db:       dbQueries,
		Commands: cmds,
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("No command provided. Available commands:")
		fmt.Println()
		fmt.Println("****")
		cmds.Run(&state, commands.Command{Name: "help"})
		fmt.Println("****")
		fmt.Println()
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := commands.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = cmds.Run(&state, cmd)
	if err != nil {
		fmt.Printf("****\nError executing command '%s'\n%v\n****\n", cmd.Name, err)
		os.Exit(1)
	}
}
