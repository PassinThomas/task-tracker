package cmd

import (
	"fmt"

	"github.com/PassinThomas/task-tracker/internal/utils"
	"strconv"

	"github.com/spf13/cobra"
)


var (
	updateCmd = &cobra.Command{
	Use:	"update",
	Short:	"Command to update to-do list",
	SilenceUsage: true,
	Args:	cobra.ExactArgs(1),
	RunE:	func(cmd *cobra.Command, args []string) error {
		utils.Debug("Run update process...",
			"cmd", "update",
			"task_id", args[0],
			"done", flg.Done,
			"new_task_name", flg.NewTitle,
		)
		
		id, errAtoi := strconv.Atoi(args[0])
		if errAtoi != nil {
			return fmt.Errorf("Fail conversion ID of updateCmd")
		}
	
		task, err := taskService.Update(id, *flg)
		if err != nil {
			return err
		}

		utils.Debug("Update task", task)
		fmt.Printf("✓ Task %d updated successfully\n", task.ID)

		return nil
	},
})
