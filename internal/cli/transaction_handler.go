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

func AnalyzeTransactions(s *State, cmd *Command) error {
	if len(cmd.Arguments) != 3 {
		return fmt.Errorf("Three Arguments needed for Analyze Transactions: x days y months z years")
	}
	daysString, monthsString, yearsString := cmd.Arguments[0], cmd.Arguments[1], cmd.Arguments[2]

	days, err := utils.StringToInt32(daysString)
	if err != nil {
		return fmt.Errorf("Error Parsing Days: %w", err)
	}

	months, err := utils.StringToInt32(monthsString)
	if err != nil {
		return fmt.Errorf("Error Parsing Months: %w", err)
	}

	years, err := utils.StringToInt32(yearsString)
	if err != nil {
		return fmt.Errorf("Error Parsing Years: %w", err)
	}

	userNumber := s.GetConfig().GetUser()
	// get user-id
	user, err := s.GetDb().GetUserByNumber(context.Background(), userNumber)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	// get Analysis
	report, err := s.GetDb().GetAnalysis(context.Background(), database.GetAnalysisParams{
		UserID: user.ID,
		Days:   days,
		Months: months,
		Years:  years,
	})
	if err != nil {
		return fmt.Errorf("Error Getting Analysis: %w", err)
	}

	log.Printf("Analysis for the user: %s for the past %d days %d months %d years \n", userNumber, days, months, years)
	for _, row := range report {
		log.Printf("LabelName: %s TotalCost: %v \n", row.Labelname, row.TotalCost)
	}

	return nil
}

func GetTransactions(s *State, cmd *Command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("No Arguments needed for Getting Transactions ")
	}
	userNumber := s.GetConfig().GetUser()

	// get user-id
	user, err := s.GetDb().GetUserByNumber(context.Background(), userNumber)
	if err != nil {
		return fmt.Errorf("Error getting user: %w", err)
	}

	transactions, err := s.GetDb().GetTransactions(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting Transactions: %w", err)
	}

	log.Printf("Transactions Fetched: \n")
	for _, transaction := range transactions {
		log.Printf("Transaction Id: %s ,Object: %s,LabelName: %s,Cost:  %f \n",
			transaction.ID,
			transaction.ObjectName,
			transaction.Labelname,
			transaction.Cost)
	}

	return nil
}
