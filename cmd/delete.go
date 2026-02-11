package cmd

import (
	"fmt"
	"strconv"

	"task/internal/service"
	"task/internal/utils"

	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
	Use:	"delete",
	Short:	"Command to delete to-do list",
	// Long: "Delete a task from the to-do list by its name. Example: task delete -d \"task name\"",
	Args: cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE:	func(cmd *cobra.Command, args []string) error {
		utils.Debug("Run delete process...",
			"cmd", "delete",
			"task_id", args[0],
		)
		v, errAtoi := strconv.Atoi(args[0])
		if errAtoi != nil {
			return fmt.Errorf("Echec conversion of delete cmd")
		}
		task, err := service.Delete(v)
		if err != nil {
			return err
		}
		utils.Debug("Task deleted", task)
		fmt.Printf("✓ Task %d deleted successfully\n", task.ID)
		return nil
	},
})
