package cmd

import (
	"fmt"

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
		err := service.Delete(args[0])
		if err != nil {
			return fmt.Errorf("Delete task impossible: %w", err)
		}
		utils.Vlog("Task deleted")
		return nil
	},
})
