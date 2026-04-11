package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Unique-GIT/expense-tracker/internal/cli"
	"github.com/Unique-GIT/expense-tracker/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	db_url := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Printf("Error Getting the database URL: %v", err)
		return
	}

	dbQueries := database.New(db)
	current_state := cli.NewState(dbQueries)

	commandRegistry := cli.NewRegistrar()

	commandRegistry.Register("add-user", cli.AddUser)
	commandRegistry.Register("get-users", cli.GetAllUsers)
	commandRegistry.Register("reset-users", cli.ResetUsers)
	commandRegistry.Register("login", cli.LoginUser)
	commandRegistry.Register("user", cli.GetCurrentUser)
	commandRegistry.Register("add-label", cli.AddLabel)

	arguments := os.Args
	if len(arguments) < 2 {
		log.Println("No Command Sent!")
		os.Exit(1)
	}

	commandName := arguments[1]
	commandArguments := arguments[2:]

	err = commandRegistry.Run(current_state, cli.NewCommand(commandName, commandArguments))
	if err != nil {
		log.Printf("Error running the command : %v", err)
		os.Exit(1)
	}
}
