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
		SilenceUsage: true,
		Args:	cobra.ExactArgs(1),
		RunE:	func(cmd *cobra.Command, args []string) error {
			utils.Debug("Run update process...",
				"cmd", "update",
				"task_id", args[0],
				"done", markDone,
				"new_task_name", newTitle,
			)
			value, errAtoi := strconv.Atoi(args[0])
			if errAtoi != nil {
				return fmt.Errorf("Argument must be a digit")
			}

			if value < 1 {
				 return fmt.Errorf("Id doesn't exist")
			}
			task, err := service.Update(value, &markDone, &newTitle)
			if err != nil {
				return err
			}
			utils.Debug("Update task", task)
			fmt.Printf("✓ Task %d updated successfully\n", task.ID)
			return nil
		},
	}
)
