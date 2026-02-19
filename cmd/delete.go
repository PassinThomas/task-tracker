package cmd

import (
	"fmt"
	"strconv"

	"github.com/PassinThomas/task-tracker/internal/utils"

	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
	Use:	"delete",
	Short:	"Command to delete to-do list",
	Args: cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE:	func(cmd *cobra.Command, args []string) error {
		utils.Debug("Run delete process...",
			"cmd", "delete",
			"task_id", args[0],
		)

		id, errAtoi := strconv.Atoi(args[0])
		if errAtoi != nil {
			return fmt.Errorf("Fail conversion ID of deleteCmd")
		}

		task, err := taskService.Delete(id)
		if err != nil {
			return err
		}

		utils.Debug("Task deleted", task)
		fmt.Printf("✓ Task %d deleted successfully\n", task.ID)
		
		return nil
	},
})
