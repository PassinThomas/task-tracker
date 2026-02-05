package cmd

import (
	"fmt"

	"task/internal/service"
	"task/internal/utils"

	"github.com/spf13/cobra"
)

var (
	delete string
	deleteCmd = &cobra.Command{
	Use:	"delete",
	Short:	"Command to delete to-do list",
	// Long: "Delete a task from the to-do list by its name. Example: task delete -d \"task name\"",
	RunE:	func(cmd *cobra.Command, args []string) error {
		err := service.Delete(delete)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		utils.Vlog(utils.Verbose, "Task deleted")
		return nil
	},
})
