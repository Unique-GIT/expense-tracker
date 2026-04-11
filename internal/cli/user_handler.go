package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/Unique-GIT/expense-tracker/internal/database"
)

func AddUser(s *State, c *Command) error {
	if len(c.Arguments) != 2 {
		return fmt.Errorf("Two arguments needed for AddUser: Name and Number")
	}

	name := c.Arguments[0]
	number := c.Arguments[1]

	user, err := s.GetDb().AddUser(
		context.Background(),
		database.AddUserParams{
			Username:   name,
			Usernumber: number,
		})
	if err != nil {
		return fmt.Errorf("Error Creating User: %w", err)
	}

	log.Printf("User Added: \n")
	log.Printf("UserName: %s \n", user.Username)
	log.Printf("UserNumber: %s \n", user.Usernumber)

	return nil
}

func GetAllUsers(s *State, c *Command) error {
	if len(c.Arguments) > 0 {
		return fmt.Errorf("No arguments are needed to Get all users")
	}

	users, err := s.GetDb().GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get all users: %w", err)
	}

	log.Printf("Users in the database: \n")
	for _, user := range users {
		log.Printf("UserName: %s UserNumber: %s \n", user.Username, user.Usernumber)
	}

	return nil
}

func ResetUsers(s *State, c *Command) error {
	if len(c.Arguments) > 0 {
		return fmt.Errorf("No arguments are needed to Get all users")
	}

	err := s.GetDb().ResetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to reset users: %w", err)
	}

	log.Printf("Reset All Users: Successfull")
	return nil
}

func LoginUser(s *State, c *Command) error {
	if len(c.Arguments) != 1 {
		return fmt.Errorf("One argument needed for Login user: Number")
	}

	number := c.Arguments[0]
	user, err := s.GetDb().GetUserByNumber(context.Background(), number)
	if err != nil {
		return fmt.Errorf("Error Logging in: %v", err)
	}

	err = s.GetConfig().SetUser(user.Username, user.Usernumber)
	if err != nil {
		return fmt.Errorf("Error Logging in: %v", err)
	}

	log.Printf("User Loggedd In: \n")
	log.Printf("UserName: %s \n", user.Username)
	log.Printf("UserNumber: %s \n", user.Usernumber)
	return nil
}

func GetCurrentUser(s *State, c *Command) error {
	if len(c.Arguments) > 0 {
		return fmt.Errorf("No arguments needed for getting current user")
	}

	user := s.GetConfig().GetUser()

	log.Printf("User Loggedd In: \n")
	log.Printf("User: %s \n", user)
	return nil
}
