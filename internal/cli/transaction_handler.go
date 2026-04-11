package cli

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Unique-GIT/expense-tracker/internal/database"
	"github.com/Unique-GIT/expense-tracker/internal/utils"
)

func AddTransaction(s *State, cmd *Command) error {
	if len(cmd.Arguments) != 3 {
		return fmt.Errorf("Three Arguments are expected for adding Transactions: LabelName, Object Name and Cost")
	}
	labelName, object, cost := cmd.Arguments[0], cmd.Arguments[1], cmd.Arguments[2]
	labelName = strings.ToLower(labelName)

	floatingCost, err := utils.ParseFloat32(cost)
	if err != nil {
		return fmt.Errorf("Error parsing cost to floating point: %w", err)
	}

	// get user-number
	userNumber := s.GetConfig().GetUser()

	// get user-id
	user, err := s.GetDb().GetUserByNumber(context.Background(), userNumber)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	//get-label-id
	label, err := s.GetDb().GetLabelId(context.Background(), database.GetLabelIdParams{
		UserID:    user.ID,
		Labelname: labelName,
	})
	if err != nil {
		return fmt.Errorf("Error Getting labelId: %w", err)
	}

	// add-transaction
	trx, err := s.GetDb().AddTransaction(context.Background(), database.AddTransactionParams{
		UserID:     user.ID,
		LabelID:    label,
		ObjectName: object,
		Cost:       floatingCost,
	})
	if err != nil {
		return fmt.Errorf("Error Adding Transaction: %w", err)
	}

	log.Printf("Added Transaction successully !\n")
	log.Printf("Transaction ID: %s \n", trx.ID)
	log.Printf("Object: %s \n", trx.ObjectName)
	log.Printf("Cost: %f \n", trx.Cost)

	return nil
}
