package cmd

import (
	"fmt"

	"task/internal/utils"
	"task/internal/service"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	markDone bool
	newTitle string
		updateCmd = &cobra.Command{
		Use:	"update",
		Short:	"Command to update to-do list",
		Args:	cobra.ExactArgs(1),
		RunE:	func(cmd *cobra.Command, args []string) error {
			value, errAtoi := strconv.Atoi(args[0])
			if errAtoi != nil {
				return fmt.Errorf("Argument must be a digit")
			} 
			if value < 1 { return fmt.Errorf("Id doesn't exist") }
			err := service.Update(value, &markDone, &newTitle)
			if err != nil {
				return fmt.Errorf("failed to update: %w", err)
			}
			utils.Vlog("Task updated")
			return nil
		},
	}
)
