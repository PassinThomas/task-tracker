package cmd

import (
	"fmt"

	"task/internal/utils"

	"github.com/spf13/cobra"
)


var addCmd = &cobra.Command{
	Use:	"add",
	Short:	"Command to add to-do list",
	Args: cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE:	func(cmd *cobra.Command, args []string) error {
		utils.Debug("Run add process...",
			"cmd", "add",
			"task_name", args[0],
		)
		errParse := utils.ParseStr(args[0])
		if errParse != nil {
			return fmt.Errorf("Bad format: %w", errParse)
		}
		task, err := taskService.Add(args[0])
		if err != nil {
			return err
		}
		utils.Debug("task added", task)
		fmt.Printf("✓ Task %d created successfully\n", task.ID)
		return nil
	},
}

