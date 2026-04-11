package cli

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Unique-GIT/expense-tracker/internal/database"
)

func AddLabel(s *State, cmd *Command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("One Argument is expected for adding Label: LabelName")
	}
	labelName := cmd.Arguments[0]
	labelName = strings.ToLower(labelName)

	// get user-number
	userNumber := s.GetConfig().GetUser()

	// get user-id
	user, err := s.GetDb().GetUserByNumber(context.Background(), userNumber)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	// add label
	label, err := s.GetDb().AddLabel(context.Background(), database.AddLabelParams{
		Labelname: labelName,
		UserID:    user.ID,
	})

	log.Printf("Label Added \n")
	log.Printf("User: %s", userNumber)
	log.Printf("LabelName: %s", label.Labelname)
	log.Printf("Label-id: %s", label.ID)
	return nil
}

func GetLabels(s *State, cmd *Command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("One Argument is expected for adding Label: LabelName")
	}

	// get user-number
	userNumber := s.GetConfig().GetUser()

	// get user-id
	user, err := s.GetDb().GetUserByNumber(context.Background(), userNumber)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	// get labels
	labels, err := s.GetDb().GetLabels(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error Getting Labels: %w", err)
	}

	log.Printf("Labels fetched: \n")
	for _, label := range labels {
		log.Printf("%s \n", label)
	}

	return nil
}
